package serialization

type OptionSerialization struct {
	QuestionId   uint   `json:"question_id"`
	OptionNumber string `json:"option_number"`
	OptionText   string `json:"option_text"`
}
