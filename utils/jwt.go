package utils

import (
	"github.com/dgrijalva/jwt-go"
	"survey_backend/global"
	"survey_backend/models"
	"time"
)

func GenerateJWT(user *models.UserModel) (string, error) {
	// 创建一个新的 JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置负载信息
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["nick_name"] = user.NickName
	claims["sex"] = user.Sex
	claims["phone"] = user.Phone
	claims["avatar"] = user.Avatar
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Exp)).Unix() // 设置过期时间，这里设置为1天

	// 使用密钥对 JWT 进行签名
	signingKey := []byte(global.Config.Jwt.Secret)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type JwtPayLoad struct {
	Avatar   string `json:"avatar"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Sex      int    `json:"sex"`
	UserId   uint   `json:"user_id"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

func VerifyJWT(tokenString string) (*CustomClaims, error) {
	// 解析 JWT
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名密钥
		signingKey := []byte(global.Config.Jwt.Secret)
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*CustomClaims)
	return claims, nil
}
