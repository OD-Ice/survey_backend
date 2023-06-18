package service

import (
	"gorm.io/gorm"
	"survey_backend/enum"
	"survey_backend/models"
)

func GetRoughlyAnswerCountService(db *gorm.DB, questionnaireId uint) int {
	var count int64
	var roughlyAnswerModel models.RoughlyAnswerModel
	db.Model(&roughlyAnswerModel).Where("questionnaire_id = ?", questionnaireId).Count(&count)
	return int(count)
}

func GetAnswerDataService(db *gorm.DB, questionnaireId uint) []models.AnswerModel {
	var answerModels []models.AnswerModel
	db.Preload("RoughlyAnswerModel", "questionnaire_id = ?", questionnaireId).Preload("QuestionModel").Find(&answerModels)
	return answerModels
}

type AnswerCount struct {
	Count       int    `json:"count"`
	QuestionID  uint   `json:"question_id"`
	OptionID    uint   `json:"option_id"`
	OptionScale string `json:"option_scale"`
	OptionText  string `json:"option_text"`
}

func GetAnswerCountService(db *gorm.DB, questionnaireId uint) []AnswerCount {
	var result []AnswerCount

	// 执行查询
	db.Table("answer_models").
		Select("count(*) as count, answer_models.question_id, answer_models.option_id, option_models.option_text").
		Joins("LEFT JOIN option_models ON option_models.id = answer_models.option_id").
		Joins("LEFT JOIN question_models ON question_models.id = answer_models.question_id").
		Joins("LEFT JOIN roughly_answer_models ON roughly_answer_models.id = answer_models.roughly_answer_model_id").
		Where("question_models.question_type IN (?,?)", enum.Single, enum.Multiple).
		Where("roughly_answer_models.questionnaire_id = ?", questionnaireId).
		Group("answer_models.question_id, answer_models.option_id").
		Scan(&result)

	// 处理查询结果
	//for _, item := range result {
	//	fmt.Printf("Count: %d, QuestionID: %d, OptionID: %d\n", item.Count, item.QuestionID, item.OptionID)
	//}
	return result
}

func GetSubjectiveAnalyzeDataService(db *gorm.DB, questionId uint, page int, results int) []models.AnswerModel {
	var answerModels []models.AnswerModel
	db.Table("answer_models").
		Joins("LEFT JOIN question_models ON question_models.id = answer_models.question_id").
		Where("answer_models.question_id = ?", questionId).
		Where("question_models.question_type = ?", enum.Subjective).
		Limit(results).Offset((page - 1) * results).
		Order("-answer_models.created_at").
		Find(&answerModels)
	return answerModels
}
