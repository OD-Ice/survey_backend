package models

type RoleModel struct {
	BaseModel
	Name           string `gorm:"size:64" json:"name"`
	PermissionList string `json:"permission_list"`
}
