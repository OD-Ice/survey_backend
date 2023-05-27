package serialization

type QuestionnaireSerialization struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	UserId      int    `json:"user_id"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
