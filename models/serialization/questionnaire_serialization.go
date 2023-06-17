package serialization

type QuestionnaireSerialization struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	UserId      uint   `json:"user_id"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type AnswerListSerialization struct {
	QuestionId   uint   `json:"question_id"`
	OptionIdList []uint `json:"option_id_list"`
	OptionId     uint   `json:"option_id"`
	TextAnswer   string `json:"text_answer"`
}

type AnswerSerialization struct {
	QuestionnaireId uint                      `json:"questionnaire_id,string"`
	AnswerList      []AnswerListSerialization `json:"answer_list"`
}
