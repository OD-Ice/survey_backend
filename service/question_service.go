package service

import (
	"survey_backend/enum"
	"survey_backend/global"
	"survey_backend/models"
)

func CreateQuestionService(questionnaireId uint, questionNum uint, questionText string, questionType int, minOption int, maxOption int) uint {
	dataModel := models.QuestionModel{
		QuestionnaireId: questionnaireId,
		QuestionNum:     questionNum,
		QuestionText:    questionText,
		QuestionType:    questionType,
		MinOption:       minOption,
		MaxOption:       maxOption,
		Status:          enum.Normal,
	}
	global.Db.Create(&dataModel)
	return dataModel.Id
}

func UpdateQuestionService(questionId uint, questionText string, questionType int) {
	var questionModel models.QuestionModel
	result := global.Db.First(&questionModel, questionId)
	result.Updates(map[string]any{
		"question_text": questionText,
		"question_type": questionType,
	})
}

func GetQuestionListService(questionnaireId uint) []models.QuestionModel {
	var questionModels []models.QuestionModel
	global.Db.Where("questionnaire_id = ?", questionnaireId).Order("question_num").Find(&questionModels)
	return questionModels
}

func GetQuestionCountService(questionnaireId uint) uint {
	var count int64
	var questionModel models.QuestionModel
	global.Db.Model(&questionModel).Where("questionnaire_id = ?", questionnaireId).Count(&count)
	return uint(count)
}
