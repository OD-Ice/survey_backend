package serialization

type QuestionnaireSerialization struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	UserId      uint   `json:"user_id"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
