package serialization

type QuestionSerialization struct {
	BaseSerialization
	QuestionnaireId uint                  `json:"questionnaire_id,string" form:"questionnaire_id"`
	QuestionNum     uint                  `json:"question_num,string"`
	QuestionText    string                `json:"question_text"`
	QuestionType    int                   `json:"question_type,string"`
	MinOption       int                   `json:"min_option,string"`
	MaxOption       int                   `json:"max_option,string"`
	OptionList      []OptionSerialization `json:"option_list"`
	Id              uint                  `json:"id" form:"id"`
}
