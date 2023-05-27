package questionnaire_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"survey_backend/enum"
	"survey_backend/global"
	"survey_backend/models"
	"survey_backend/models/res"
	"survey_backend/models/serialization"
	"time"
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

	global.Db.Create(&models.QuestionnaireModel{
		Title:       requestBody.Title,
		Description: requestBody.Description,
		UserId:      currUser.Id,
		Status:      enum.Normal,
	})
	res.OkWith(c)
}

// UpdateQuestionnaireView 修改调查问卷
func (QuestionnaireApi) UpdateQuestionnaireView(c *gin.Context) {
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var questionnaireModel models.QuestionnaireModel
	// 获取要更新的那条数据
	result := global.Db.First(&questionnaireModel, requestBody.Id)

	//result.Updates(&models.QuestionnaireModel{
	//	Title:       requestBody.Name,
	//	Description: requestBody.Description,
	//})
	result.Updates(map[string]any{
		"title":       requestBody.Title,
		"description": requestBody.Description,
	})
	res.OkWith(c)
}

// GetQuestionnaireView 查询用户的调查问卷列表
func (QuestionnaireApi) GetQuestionnaireView(c *gin.Context) {
	// 拿到当前用户
	_currUser, _ := c.Get("currUser")
	currUser := _currUser.(*models.UserModel)
	// 查询
	var questionnaireModels []models.QuestionnaireModel
	global.Db.Where("user_id = ?", currUser.Id).Select("id, title, user_id, description, status").Find(&questionnaireModels)

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
	var questionnaireModel models.QuestionnaireModel
	result := global.Db.First(&questionnaireModel, requestBody.Id)
	result.Updates(map[string]any{
		"deleted_at": time.Now(),
		"status":     enum.Deleted,
	})
	//global.Db.Delete(&questionnaireModel, "id = ?", requestBody.Id)
	res.OkWith(c)
}
