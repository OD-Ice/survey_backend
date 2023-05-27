package questionnaire_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"survey_backend/models"
	"survey_backend/models/res"
	"survey_backend/models/serialization"
	"survey_backend/service"
)

// CreateQuestionnaireView 创建调查问卷
func (QuestionnaireApi) CreateQuestionnaireView(c *gin.Context) {
	// 捕捉报错
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered:", r)
	//	}
	//}()

	var requestBody serialization.QuestionnaireSerialization
	_currUser, _ := c.Get("currUser")
	currUser := _currUser.(*models.UserModel)
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		fmt.Println(err.Error())
		res.FailWithCode(res.ParameterError, c)
		return
	}

	questionnaireId := service.CreateQuestionnaireService(requestBody.Title, requestBody.Description, currUser.Id)
	data := map[string]any{"questionnaireId": questionnaireId}
	res.OkWithData(data, c)
}

// UpdateQuestionnaireView 修改调查问卷
func (QuestionnaireApi) UpdateQuestionnaireView(c *gin.Context) {
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	service.UpdateQuestionnaireService(requestBody.Id, requestBody.Title, requestBody.Description)

	res.OkWith(c)
}

// GetQuestionnaireView 查询用户的调查问卷列表
func (QuestionnaireApi) GetQuestionnaireView(c *gin.Context) {
	// 拿到当前用户
	_currUser, _ := c.Get("currUser")
	currUser := _currUser.(*models.UserModel)
	// 查询
	questionnaireModels := service.GetQuestionnaireByUserIdService(currUser.Id)

	var data []map[string]any
	for _, item := range questionnaireModels {
		data = append(
			data, map[string]any{
				"id":          item.Id,
				"title":       item.Title,
				"user_id":     item.UserId,
				"description": item.Description,
				"status":      item.Status,
			},
		)
	}

	res.OkWithData(data, c)
}

// DeleteQuestionnaireView 删除调查问卷
func (QuestionnaireApi) DeleteQuestionnaireView(c *gin.Context) {
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		fmt.Println(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}
	service.DelQuestionnaireService(requestBody.Id)

	res.OkWith(c)
}
