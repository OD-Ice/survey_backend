package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"net/http"
	"survey_backend/global"
	"survey_backend/models"
	"survey_backend/service"
	"survey_backend/utils"
)

// AuthMiddleware 中间件函数，用于从请求中获取当前用户的信息并将其保存到上下文中
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设在请求头中使用了身份验证令牌（例如 JWT），并将用户信息存储在令牌中

		// 从请求头中获取身份验证令牌
		token := c.GetHeader("token")

		// 验证是否注销
		result, err := global.Redis.Get(fmt.Sprintf("logout_%s", token)).Result()
		if err != redis.Nil && result == "1" {
			// key存在，已注销
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户已注销"})
			return
		}

		// 根据令牌进行用户认证，并获取当前用户的信息
		user, exp, err := authenticateUser(c, token)
		if err != nil {
			// 认证失败，返回未经授权的错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}

		// 将当前用户的信息保存到上下文中，以便后续处理程序使用
		c.Set("currUser", user)
		c.Set("exp", exp)

		// 继续处理后续的请求
		c.Next()
	}
}

// 模拟用户认证的函数，根据令牌获取用户信息
func authenticateUser(c *gin.Context, token string) (*models.UserModel, int64, error) {
	db, _ := c.Get("db")
	// 在这里进行身份验证和授权的逻辑，根据令牌获取用户信息
	jwt, err := utils.VerifyJWT(token)
	if err != nil {
		return nil, 0, err
	}
	userId := jwt.JwtPayLoad.UserId
	// 验证用户是否存在
	user, ok := service.GetUserById(db.(*gorm.DB), userId)
	if !ok {
		panic("用户不存在")
	}

	return user, jwt.ExpiresAt, nil
}

func ReqHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 设置允许跨域请求的域名
	}
}

func ApiMySql() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := global.Db.Begin()
		defer func() {
			if r := recover(); r != nil {
				// 发生 panic，回滚事务
				db.Rollback()
				panic(r)
			} else if c.IsAborted() {
				// 请求被中断，回滚事务
				db.Rollback()
			} else {
				// 提交事务
				db.Commit()
			}
		}()
		c.Set("db", db)
		c.Next()
	}
}
