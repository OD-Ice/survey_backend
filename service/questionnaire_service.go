package service

import (
	"survey_backend/enum"
	"survey_backend/global"
	"survey_backend/models"
	"time"
)

func CreateQuestionnaireService(title string, description string, userId uint) uint {
	dataModel := models.QuestionnaireModel{
		Title:       title,
		Description: description,
		UserId:      userId,
		Status:      enum.Normal,
	}
	global.Db.Create(&dataModel)
	return dataModel.Id
}

func UpdateQuestionnaireService(questionnaireId uint, title string, description string) {
	var questionnaireModel models.QuestionnaireModel
	result := global.Db.First(&questionnaireModel, questionnaireId)

	//result.Updates(&models.QuestionnaireModel{
	//	Title:       requestBody.Name,
	//	Description: requestBody.Description,
	//})
	result.Updates(map[string]any{
		"title":       title,
		"description": description,
	})
}

func GetQuestionnaireByUserIdService(userId uint, page int, results int) []models.QuestionnaireModel {
	var questionnaireModels []models.QuestionnaireModel
	global.Db.Where("user_id = ?", userId).Limit(results).Offset((page - 1) * results).Order("-created_at").Find(&questionnaireModels)
	return questionnaireModels
}

func DelQuestionnaireService(questionnaireId uint) {
	var questionnaireModel models.QuestionnaireModel
	result := global.Db.First(&questionnaireModel, questionnaireId)
	result.Updates(map[string]any{
		"deleted_at": time.Now(),
		"status":     enum.Deleted,
	})
	// global.Db.Delete(&questionnaireModel, "id = ?", requestBody.Id)
}
