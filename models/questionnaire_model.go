package models

// QuestionnaireModel 问卷表
type QuestionnaireModel struct {
	BaseModel
	Title       string    `json:"title"`
	UserModel   UserModel `gorm:"foreignKey:UserId" json:"user_model"`
	UserId      uint      `json:"user_id"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
}

// QuestionModel 问题表
type QuestionModel struct {
	BaseModel
	QuestionnaireModel QuestionnaireModel `gorm:"foreignKey:QuestionnaireId" json:"questionnaire_model"`
	QuestionnaireId    uint               `json:"questionnaire_id"`
	QuestionNum        uint               `json:"question_num"` // 问题序号，用来排序
	QuestionText       string             `json:"question_text"`
	QuestionType       int                `json:"question_type"` // 问题类型（单选、多选、文本）
	MinOption          int                `json:"min_option"`    // 最少选项数
	MaxOption          int                `json:"max_option"`    // 最多选项数
	Status             int                `json:"status"`
}

// OptionModel 选项表
type OptionModel struct {
	BaseModel
	QuestionModel QuestionModel `gorm:"foreignKey:QuestionId" json:"question_model"`
	QuestionId    uint          `json:"question_id"`
	OptionNumber  string        `json:"option_number"` // 选项编号(a,b,c,d...)
	OptionText    string        `json:"option_text"`
	Status        int           `json:"status"`
}
