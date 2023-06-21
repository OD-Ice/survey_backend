package routers

import (
	"survey_backend/api"
	"survey_backend/middleware"
)

func (router RouterGroup) QuestionnaireRouter() {
	// 分组
	QuestionnaireRouter := router.Group("questionnaire")
	// 判断用户是否登录的中间件
	QuestionnaireRouter.Use(middleware.AuthMiddleware())
	QuestionnaireApi := api.ApiGroupApp.QuestionnaireApi
	QuestionnaireRouter.POST("/create", QuestionnaireApi.CreateQuestionnaireView)
	QuestionnaireRouter.POST("/update", QuestionnaireApi.UpdateQuestionnaireView)
	QuestionnaireRouter.GET("/get_list", QuestionnaireApi.GetQuestionnaireView)
	QuestionnaireRouter.POST("/delete", QuestionnaireApi.DeleteQuestionnaireView)
	QuestionnaireRouter.POST("/edit_status", QuestionnaireApi.EditStatusView)
	QuestionnaireRouter.POST("/answer", QuestionnaireApi.AnswerView)
	QuestionnaireRouter.GET("/analyze", QuestionnaireApi.GetAnalyzeDataView)
	QuestionnaireRouter.GET("/answer_count", QuestionnaireApi.GetRoughlyAnswerCountView)
	QuestionnaireRouter.GET("/get", QuestionnaireApi.GetQuestionnaireByIdView)
	QuestionnaireRouter.GET("/get_subjective_analyze", QuestionnaireApi.GetSubjectiveAnalyzeDataView)
}
