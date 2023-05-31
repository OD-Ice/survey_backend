package serialization

type BaseSerialization struct {
	Page    int `form:"page"`
	Results int `form:"results"`
}
