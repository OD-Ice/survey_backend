package questionnaire_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"survey_backend/models"
	"survey_backend/models/res"
	"survey_backend/models/serialization"
	"survey_backend/service"
	"survey_backend/utils"
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
	db, _ := c.Get("db")
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		fmt.Println(err.Error())
		res.FailWithCode(res.ParameterError, c)
		return
	}

	questionnaireId := service.CreateQuestionnaireService(db.(*gorm.DB), requestBody.Title, requestBody.Description, currUser.Id)
	data := map[string]any{"questionnaireId": questionnaireId}
	res.OkWithData(data, c)
}

// UpdateQuestionnaireView 修改调查问卷
func (QuestionnaireApi) UpdateQuestionnaireView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	service.UpdateQuestionnaireService(db.(*gorm.DB), requestBody.Id, requestBody.Title, requestBody.Description)

	res.OkWith(c)
}

// GetQuestionnaireView 查询用户的调查问卷列表
func (QuestionnaireApi) GetQuestionnaireView(c *gin.Context) {
	// 拿到当前用户
	_currUser, _ := c.Get("currUser")
	db, _ := c.Get("db")
	currUser := _currUser.(*models.UserModel)
	// 查询
	var requestBody serialization.BaseSerialization
	err := c.ShouldBindQuery(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	questionnaireModels := service.GetQuestionnaireByUserIdService(db.(*gorm.DB), currUser.Id, requestBody.Page, requestBody.Results)

	var data []map[string]any
	for _, item := range questionnaireModels {
		data = append(
			data, map[string]any{
				"id":          item.Id,
				"title":       item.Title,
				"user_id":     item.UserId,
				"description": item.Description,
				"status":      item.Status,
				"create_at":   item.CreatedAt.Format("2006-01-02 15:04:05"),
				"update_at":   item.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		)
	}

	res.OkWithData(data, c)
}

// DeleteQuestionnaireView 删除调查问卷
func (QuestionnaireApi) DeleteQuestionnaireView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	service.DelQuestionnaireService(db.(*gorm.DB), requestBody.Id)

	res.OkWith(c)
}

// EditStatusView 编辑问卷状态
func (QuestionnaireApi) EditStatusView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	service.UpdateQuestionnaireStatusService(db.(*gorm.DB), requestBody.Id, requestBody.Status)
	res.OkWith(c)
}

// AnswerView 回答问卷
func (QuestionnaireApi) AnswerView(c *gin.Context) {
	_currUser, _ := c.Get("currUser")
	db, _ := c.Get("db")
	currUser := _currUser.(*models.UserModel)
	var requestBody serialization.AnswerSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	// 验证问卷状态
	questionnaire := service.GetPublishedQuestionnaireByIdService(db.(*gorm.DB), requestBody.QuestionnaireId)
	if questionnaire == nil {
		res.FailWithMsg("问卷不存在或没有发布", c)
		return
	}
	// 创建答卷
	roughlyAnswerId := service.CreateRoughlyAnswerService(db.(*gorm.DB), questionnaire.Id, currUser.Id)
	// 创建回答
	var answerList []models.AnswerModel
	for _, answer := range requestBody.AnswerList {
		if answer.OptionId != 0 {
			// 单选
			answerList = append(answerList, models.AnswerModel{
				RoughlyAnswerModelId: roughlyAnswerId,
				QuestionId:           answer.QuestionId,
				OptionId:             answer.OptionId,
			})
		} else if utils.IsUintListEmpty(answer.OptionIdList) {
			// 简答
			answerList = append(answerList, models.AnswerModel{
				RoughlyAnswerModelId: roughlyAnswerId,
				QuestionId:           answer.QuestionId,
				TextAnswer:           answer.TextAnswer,
			})
		} else {
			// 多选
			for _, optionId := range answer.OptionIdList {
				answerList = append(answerList, models.AnswerModel{
					RoughlyAnswerModelId: roughlyAnswerId,
					QuestionId:           answer.QuestionId,
					OptionId:             optionId,
				})
			}
		}
	}
	service.BatchCreateAnswerService(db.(*gorm.DB), answerList)
	res.OkWith(c)
}

func (QuestionnaireApi) GetQuestionnaireByIdView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindQuery(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	questionnaire := service.GetQuestionnaireByIdService(db.(*gorm.DB), requestBody.Id)
	res.OkWithData(questionnaire, c)
}
