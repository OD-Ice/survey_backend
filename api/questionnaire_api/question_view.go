package questionnaire_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"survey_backend/enum"
	"survey_backend/models"
	"survey_backend/models/res"
	"survey_backend/models/serialization"
	"survey_backend/service"
	"survey_backend/utils"
)

// CreateQuestionView 添加问题
func (QuestionApi) CreateQuestionView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// 查询已有问题的数量
	questionCount := service.GetQuestionCountService(db.(*gorm.DB), requestBody.QuestionnaireId)
	// 创建问题
	questionId := service.CreateQuestionService(
		db.(*gorm.DB),
		requestBody.QuestionnaireId,
		questionCount+1,
		requestBody.QuestionText,
		requestBody.QuestionType,
		requestBody.MinOption,
		requestBody.MaxOption,
	)

	// 选择题需要创建选项
	if utils.InList(requestBody.QuestionType, []any{enum.Single, enum.Multiple}) {
		service.BatchCreateOptionService(db.(*gorm.DB), questionId, requestBody.OptionList)
	}

	data := map[string]any{"questionId": questionId}
	res.OkWithData(data, c)
}

// EditQuestionView 编辑问题
func (QuestionApi) EditQuestionView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// 修改问题
	service.UpdateQuestionService(
		db.(*gorm.DB),
		requestBody.Id,
		requestBody.QuestionText,
		requestBody.QuestionType,
	)

	// 删除已有选项
	service.DelOptionServiceByQuestion(db.(*gorm.DB), requestBody.Id)
	// 选择题需要创建选项
	if utils.InList(requestBody.QuestionType, []any{enum.Single, enum.Multiple}) {
		service.BatchCreateOptionService(db.(*gorm.DB), requestBody.Id, requestBody.OptionList)
	}

	data := map[string]any{"questionId": requestBody.Id}
	res.OkWithData(data, c)
}

// GetQuestionListView 获取问题
func (QuestionApi) GetQuestionListView(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.QuestionSerialization
	err := c.ShouldBindQuery(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// 查询
	questionList := service.GetQuestionListService(db.(*gorm.DB), requestBody.QuestionnaireId)
	var questionListData []map[string]any
	var questionIdList []uint
	for _, question := range questionList {
		questionData := map[string]any{
			"id":            question.Id,
			"question_num":  question.QuestionNum,
			"question_text": question.QuestionText,
			"question_type": question.QuestionType,
			"min_option":    question.MinOption,
			"max_option":    question.MaxOption,
		}
		questionListData = append(questionListData, questionData)
		questionIdList = append(questionIdList, question.Id)
	}
	// 查询选项
	optionList := service.GetOptionListService(db.(*gorm.DB), questionIdList)
	optionDict := map[uint][]models.OptionModel{}
	for _, option := range optionList {
		optionDict[option.QuestionId] = append(optionDict[option.QuestionId], option)
	}

	for _, question := range questionListData {
		if option, ok := optionDict[question["id"].(uint)]; ok {
			question["option_list"] = option
		}
	}

	data := map[string]any{
		"question_list":  questionListData,
		"question_count": len(questionList),
	}
	res.OkWithData(data, c)
}
