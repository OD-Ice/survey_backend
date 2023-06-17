package service

import (
	"gorm.io/gorm"
	"survey_backend/enum"
	"survey_backend/models"
)

func CreateQuestionService(db *gorm.DB, questionnaireId uint, questionNum uint, questionText string, questionType int, minOption int, maxOption int) uint {
	dataModel := models.QuestionModel{
		QuestionnaireId: questionnaireId,
		QuestionNum:     questionNum,
		QuestionText:    questionText,
		QuestionType:    questionType,
		MinOption:       minOption,
		MaxOption:       maxOption,
		Status:          enum.Normal,
	}
	db.Create(&dataModel)
	return dataModel.Id
}

func UpdateQuestionService(db *gorm.DB, questionId uint, questionText string, questionType int) {
	var questionModel models.QuestionModel
	result := db.First(&questionModel, questionId)
	result.Updates(map[string]any{
		"question_text": questionText,
		"question_type": questionType,
	})
}

func GetQuestionListService(db *gorm.DB, questionnaireId uint) []models.QuestionModel {
	var questionModels []models.QuestionModel
	db.Where("questionnaire_id = ?", questionnaireId).Order("question_num").Find(&questionModels)
	return questionModels
}

func GetQuestionCountService(db *gorm.DB, questionnaireId uint) uint {
	var count int64
	var questionModel models.QuestionModel
	db.Model(&questionModel).Where("questionnaire_id = ?", questionnaireId).Count(&count)
	return uint(count)
}
