package serialization

type LoginSerialization struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type RegisterSerialization struct {
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	NickName        string `json:"nick_name"`
	Sex             string `json:"sex"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
}
