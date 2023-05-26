package models

type AnswerModel struct {
	BaseModel
	QuestionnaireModel QuestionnaireModel `gorm:"foreignKey:QuestionnaireId" json:"questionnaire_model"`
	QuestionnaireId    int                `json:"questionnaire_id"`
	QuestionModel      QuestionModel      `gorm:"foreignKey:QuestionId" json:"question_model"`
	QuestionId         int                `json:"question_id"`
	OptionIds          string             `json:"option_ids"`
	TextAnswer         string             `json:"text_answer"`
	UserModel          UserModel          `gorm:"foreignKey:UserId" json:"user_model"`
	UserId             int                `json:"user_id"`
}
