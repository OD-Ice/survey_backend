package routers

import (
	"survey_backend/api"
	"survey_backend/middleware"
)

func (router RouterGroup) QuestionRouter() {
	// 分组
	QuestionRouter := router.Group("question")
	// 判断用户是否登录的中间件
	QuestionRouter.Use(middleware.AuthMiddleware())
	QuestionApi := api.ApiGroupApp.QuestionApi
	QuestionRouter.POST("/create", QuestionApi.CreateQuestionView)
	QuestionRouter.POST("/edit", QuestionApi.EditQuestionView)
	QuestionRouter.GET("/get_list", QuestionApi.GetQuestionListView)
}
