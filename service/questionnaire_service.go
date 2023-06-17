package service

import (
	"gorm.io/gorm"
	"survey_backend/enum"
	"survey_backend/models"
	"time"
)

func CreateQuestionnaireService(db *gorm.DB, title string, description string, userId uint) uint {
	dataModel := models.QuestionnaireModel{
		Title:       title,
		Description: description,
		UserId:      userId,
		Status:      enum.Normal,
	}
	db.Create(&dataModel)
	return dataModel.Id
}

func UpdateQuestionnaireService(db *gorm.DB, questionnaireId uint, title string, description string) {
	var questionnaireModel models.QuestionnaireModel
	result := db.First(&questionnaireModel, questionnaireId)

	//result.Updates(&models.QuestionnaireModel{
	//	Title:       requestBody.Name,
	//	Description: requestBody.Description,
	//})
	result.Updates(map[string]any{
		"title":       title,
		"description": description,
	})
}

func UpdateQuestionnaireStatusService(db *gorm.DB, questionnaireId uint, status int) {
	var questionnaireModel models.QuestionnaireModel
	result := db.First(&questionnaireModel, questionnaireId)
	if status != enum.Deleted {
		result.Updates(map[string]any{
			"status":     status,
			"deleted_at": nil,
		})
	} else {
		result.Updates(map[string]any{
			"status":     status,
			"deleted_at": time.Now(),
		})
	}
}

func GetQuestionnaireByIdService(db *gorm.DB, id uint) *models.QuestionnaireModel {
	var questionnaireModel models.QuestionnaireModel
	err := db.Where("id = ? AND status = ?", id, enum.Published).Take(&questionnaireModel).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return &questionnaireModel
}

func GetQuestionnaireByUserIdService(db *gorm.DB, userId uint, page int, results int) []models.QuestionnaireModel {
	var questionnaireModels []models.QuestionnaireModel
	db.Where("user_id = ?", userId).Limit(results).Offset((page - 1) * results).Order("-created_at").Find(&questionnaireModels)
	return questionnaireModels
}

func DelQuestionnaireService(db *gorm.DB, questionnaireId uint) {
	var questionnaireModel models.QuestionnaireModel
	result := db.First(&questionnaireModel, questionnaireId)
	result.Updates(map[string]any{
		"deleted_at": time.Now(),
		"status":     enum.Deleted,
	})
	// global.Db.Delete(&questionnaireModel, "id = ?", requestBody.Id)
}

func CreateRoughlyAnswerService(db *gorm.DB, questionnaireId uint, userId uint) uint {
	roughlyAnswerModel := models.RoughlyAnswerModel{
		QuestionnaireId: questionnaireId,
		UserId:          userId,
	}
	db.Create(&roughlyAnswerModel)
	return roughlyAnswerModel.Id
}

func BatchCreateAnswerService(db *gorm.DB, answerList []models.AnswerModel) {
	db.CreateInBatches(answerList, len(answerList))
}
