package service

import (
	"gorm.io/gorm"
	"strconv"
	"survey_backend/enum"
	"survey_backend/models"
	"survey_backend/models/serialization"
	"survey_backend/utils"
)

func CreateUser(db *gorm.DB, user *serialization.RegisterSerialization) (uint, error) {
	password, err := utils.EncryptPwd(user.Password)
	if err != nil {
		return 0, err
	}
	sex, _ := strconv.Atoi(user.Sex)
	dataModel := models.UserModel{
		UserName: user.UserName,
		Password: password,
		NickName: user.NickName,
		Avatar:   "https://img1.baidu.com/it/u=1592760890,4076657703&fm=253&fmt=auto&app=138&f=JPEG",
		Sex:      sex,
		Phone:    user.Phone,
		Email:    user.Email,
		RoleId:   1,
		Status:   enum.Normal,
	}
	db.Create(&dataModel)
	return dataModel.Id, nil
}

// GetUserByParams 通过用户名/手机号/邮箱获取用户
func GetUserByParams(db *gorm.DB, username string) (*models.UserModel, bool) {
	var userModel models.UserModel
	err := db.Take(&userModel, "user_name = ? or phone = ? or email = ?", username, username, username).Error
	if err != nil {
		// 没找到
		return nil, false
	}
	return &userModel, true
}

func GetUserById(db *gorm.DB, userId uint) (*models.UserModel, bool) {
	var userModel models.UserModel
	err := db.Take(&userModel, "id = ?", userId).Error
	if err != nil {
		// 没找到
		return nil, false
	}
	return &userModel, true
}
