package models

type AnswerModel struct {
	BaseModel
	QuestionnaireModel QuestionnaireModel `gorm:"foreignKey:QuestionnaireId" json:"questionnaire_model"`
	QuestionnaireId    uint               `json:"questionnaire_id"`
	QuestionModel      QuestionModel      `gorm:"foreignKey:QuestionId" json:"question_model"`
	QuestionId         uint               `json:"question_id"`
	OptionIds          string             `json:"option_ids"`
	TextAnswer         string             `json:"text_answer"`
	UserModel          UserModel          `gorm:"foreignKey:UserId" json:"user_model"`
	UserId             uint               `json:"user_id"`
}
