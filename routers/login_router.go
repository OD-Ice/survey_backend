package routers

import "survey_backend/api/login_api"

func (router RouterGroup) LoginRouter() {
	router.POST("/login", login_api.LoginApi{}.Login)
	router.POST("/register", login_api.LoginApi{}.Register)
}
