package questionnaire_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"survey_backend/global"
	"survey_backend/models"
	"survey_backend/models/res"
	"time"
)

type QuestionnaireSerialization struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateQuestionnaireView 创建调查问卷
func (QuestionnaireApi) CreateQuestionnaireView(c *gin.Context) {
	// 捕捉报错
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered:", r)
	//	}
	//}()

	var requestBody QuestionnaireSerialization
	_currUser, _ := c.Get("currUser")
	currUser := _currUser.(*models.UserModel)
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		fmt.Println(err.Error())
		res.FailWithCode(res.ParameterError, c)
		return
	}

	global.Db.Create(&models.QuestionnaireModel{
		Title:       requestBody.Name,
		Description: requestBody.Description,
		UserId:      currUser.Id,
	})
	res.OkWith(c)
}

// UpdateQuestionnaireView 修改调查问卷
func (QuestionnaireApi) UpdateQuestionnaireView(c *gin.Context) {
	var requestBody QuestionnaireSerialization
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
		"title":       requestBody.Name,
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
	global.Db.Where("user_id = ?", currUser.Id).Find(&questionnaireModels)

	res.OkWithData(questionnaireModels, c)
}

// DeleteQuestionnaireView 删除调查问卷
func (QuestionnaireApi) DeleteQuestionnaireView(c *gin.Context) {
	var requestBody QuestionnaireSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		fmt.Println(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var questionnaireModel models.QuestionnaireModel
	reuslt := global.Db.First(&questionnaireModel, requestBody.Id)
	reuslt.Updates(map[string]any{
		"deleted_at": time.Now(),
		"status":     -1,
	})
	//global.Db.Delete(&questionnaireModel, "id = ?", requestBody.Id)
	res.OkWith(c)
}
