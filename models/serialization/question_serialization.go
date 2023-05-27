package serialization

type QuestionSerialization struct {
	QuestionnaireId uint                  `json:"questionnaire_id"`
	QuestionNum     uint                  `json:"question_num"`
	QuestionText    string                `json:"question_text"`
	QuestionType    int                   `json:"question_type"`
	MinOption       int                   `json:"min_option"`
	MaxOption       int                   `json:"max_option"`
	OptionList      []OptionSerialization `json:"option_list"`
}
