package questionnaire_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"survey_backend/enum"
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
		fmt.Println(err.Error())
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
