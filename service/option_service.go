package service

import (
	"gorm.io/gorm"
	"survey_backend/enum"
	"survey_backend/models"
	"survey_backend/models/serialization"
)

func DelOptionServiceByQuestion(db *gorm.DB, questionId uint) {
	var dataModel models.OptionModel
	db.Where("question_id = ?", questionId).Delete(&dataModel)
}

func BatchCreateOptionService(db *gorm.DB, questionId uint, optionList []serialization.OptionSerialization) {
	var dataModels []models.OptionModel
	for _, item := range optionList {
		dataModels = append(dataModels, models.OptionModel{
			QuestionId:   questionId,
			OptionNumber: item.OptionNumber,
			OptionText:   item.OptionText,
			Status:       enum.Normal,
		})
	}
	db.CreateInBatches(dataModels, len(dataModels))
}

func GetOptionListService(db *gorm.DB, questionIdList []uint) []models.OptionModel {
	var optionModels []models.OptionModel
	db.Where("question_id in ?", questionIdList).Order("option_number").Find(&optionModels)
	return optionModels
}
