package questionnaire_api

import (
	"github.com/gin-gonic/gin"
	"survey_backend/enum"
	"survey_backend/models"
	"survey_backend/models/res"
	"survey_backend/models/serialization"
	"survey_backend/service"
	"survey_backend/utils"
)

// CreateQuestionView 添加问题
func (QuestionApi) CreateQuestionView(c *gin.Context) {
	var requestBody serialization.QuestionSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// 创建问题
	questionId := service.CreateQuestionService(
		requestBody.QuestionnaireId,
		requestBody.QuestionNum,
		requestBody.QuestionText,
		requestBody.QuestionType,
		requestBody.MinOption,
		requestBody.MaxOption,
	)

	// 选择题需要创建选项
	if utils.InList(requestBody.QuestionType, []any{enum.Single, enum.Multiple}) {
		service.BatchCreateOptionService(questionId, requestBody.OptionList)
	}

	data := map[string]any{"questionId": questionId}
	res.OkWithData(data, c)
}

// GetQuestionListView 获取问题
func (QuestionApi) GetQuestionListView(c *gin.Context) {
	var requestBody serialization.QuestionSerialization
	err := c.ShouldBindQuery(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	// 查询
	questionList := service.GetQuestionListService(requestBody.QuestionnaireId)
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
	optionList := service.GetOptionListService(questionIdList)
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
