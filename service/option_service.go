package service

import (
	"survey_backend/enum"
	"survey_backend/global"
	"survey_backend/models"
	"survey_backend/models/serialization"
)

func CreateOptionService(questionId uint, optionNumber string, optionText string) uint {
	dataModel := models.OptionModel{
		QuestionId:   questionId,
		OptionNumber: optionNumber,
		OptionText:   optionText,
		Status:       enum.Normal,
	}
	global.Db.Create(&dataModel)
	return dataModel.Id
}

func BatchCreateOptionService(questionId uint, optionList []serialization.OptionSerialization) {
	var dataModels []models.OptionModel
	for _, item := range optionList {
		dataModels = append(dataModels, models.OptionModel{
			QuestionId:   questionId,
			OptionNumber: item.OptionNumber,
			OptionText:   item.OptionText,
			Status:       enum.Normal,
		})
	}
	global.Db.CreateInBatches(dataModels, len(dataModels))
}