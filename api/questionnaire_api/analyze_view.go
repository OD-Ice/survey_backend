package questionnaire_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"survey_backend/models"
	"survey_backend/models/res"
	"survey_backend/models/serialization"
	"survey_backend/service"
)

// GetAnalyzeDataView 获取分析数据
func (QuestionnaireApi) GetAnalyzeDataView(c *gin.Context) {
	db, _ := c.Get("db")
	_currUser, _ := c.Get("currUser")
	currUser := _currUser.(*models.UserModel)
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindQuery(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	questionnaireId := requestBody.Id
	// 获取问卷
	questionnaire := service.GetQuestionnaireByIdService(db.(*gorm.DB), questionnaireId)
	if questionnaire == nil || questionnaire.UserId != currUser.Id {
		res.FailWithMsg("只有问卷的创建者可以查看统计结果", c)
		return
	}
	// 获取答卷数量
	roughlyAnswerCount := service.GetRoughlyAnswerCountService(db.(*gorm.DB), questionnaireId)

	// 获取单选和多选答案数据
	answerCountData := service.GetAnswerCountService(db.(*gorm.DB), questionnaireId)
	result := make(map[uint][]service.AnswerCount)
	for i, answer := range answerCountData {
		answerCountData[i].OptionScale = fmt.Sprintf("%.4f", float32(answer.Count)/float32(roughlyAnswerCount))
		result[answer.QuestionID] = append(result[answer.QuestionID], answerCountData[i])
	}
	res.OkWithData(result, c)
}

// GetRoughlyAnswerCountView 获取答卷数量
func (QuestionnaireApi) GetRoughlyAnswerCountView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionnaireSerialization
	err := c.ShouldBindQuery(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	questionnaireId := requestBody.Id
	roughlyAnswerCount := service.GetRoughlyAnswerCountService(db.(*gorm.DB), questionnaireId)
	res.OkWithData(map[string]int{"count": roughlyAnswerCount}, c)
}

// GetSubjectiveAnalyzeDataView 查询简答题所有答题数据
func (QuestionnaireApi) GetSubjectiveAnalyzeDataView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionSerialization
	err := c.ShouldBindQuery(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	answerList := service.GetSubjectiveAnalyzeDataService(db.(*gorm.DB), requestBody.Id, requestBody.Page, requestBody.Results)
	res.OkWithData(answerList, c)
}
