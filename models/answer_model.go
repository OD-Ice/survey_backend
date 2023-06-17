package models

type RoughlyAnswerModel struct {
	BaseModel
	UserModel          UserModel          `gorm:"foreignKey:UserId" json:"user_model"`
	UserId             uint               `json:"user_id"`
	QuestionnaireModel QuestionnaireModel `gorm:"foreignKey:QuestionnaireId" json:"questionnaire_model"`
	QuestionnaireId    uint               `json:"questionnaire_id"`
}

type AnswerModel struct {
	BaseModel
	RoughlyAnswerModel   RoughlyAnswerModel `gorm:"foreignKey:RoughlyAnswerModelId" json:"roughly_answer_model"`
	RoughlyAnswerModelId uint               `json:"roughly_answer_model_id"`
	QuestionModel        QuestionModel      `gorm:"foreignKey:QuestionId" json:"question_model"`
	QuestionId           uint               `json:"question_id"`
	OptionId             uint               `json:"option_id"`
	TextAnswer           string             `json:"text_answer"`
}
