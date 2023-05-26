package models

type UserModel struct {
	BaseModel
	UserName string `gorm:"size:64" json:"user_name"`
	Password string `gorm:"size:128" json:"password"`
	NickName string `gorm:"size:64" json:"nick_name"`
	Sex      int    `json:"sex"`
	Phone    string `gorm:"size:11" json:"phone"`
	Email    string `gorm:"size:64" json:"email"`
	RoleId   int    `json:"roleId"`
	Status   int    `json:"status"`
}
