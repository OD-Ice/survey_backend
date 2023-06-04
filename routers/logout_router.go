package routers

import (
	"survey_backend/api/logout_api"
	"survey_backend/middleware"
)

func (router RouterGroup) LogoutRouter() {
	authRouter := router.Group("auth")
	// 判断用户是否登录的中间件
	authRouter.Use(middleware.AuthMiddleware())
	authRouter.POST("/logout", logout_api.LogoutApi{}.LogOut)
}
