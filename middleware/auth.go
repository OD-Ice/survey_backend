package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"survey_backend/models"
)

// AuthMiddleware 中间件函数，用于从请求中获取当前用户的信息并将其保存到上下文中
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 假设在请求头中使用了身份验证令牌（例如 JWT），并将用户信息存储在令牌中

		// 从请求头中获取身份验证令牌
		token := c.GetHeader("Authorization")

		// 根据令牌进行用户认证，并获取当前用户的信息
		user, err := authenticateUser(token)
		if err != nil {
			// 认证失败，返回未经授权的错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// 将当前用户的信息保存到上下文中，以便后续处理程序使用
		c.Set("currUser", user)

		// 继续处理后续的请求
		c.Next()
	}
}

// 模拟用户认证的函数，根据令牌获取用户信息
func authenticateUser(token string) (*models.UserModel, error) {
	// 在这里进行身份验证和授权的逻辑，根据令牌获取用户信息
	// 这里只是一个示例，返回固定的用户信息

	// 假设令牌的格式是 "Bearer <token>"
	// 去掉前缀 "Bearer "，获取实际的令牌值
	//actualToken := token[7:]

	// 假设用户信息是从令牌中解析得到的
	user := models.UserModel{
		BaseModel: models.BaseModel{Id: 1},
		UserName:  "测试",
	}

	return &user, nil
}
