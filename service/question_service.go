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
