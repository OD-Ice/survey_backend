package login_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"survey_backend/global"
	"survey_backend/models/res"
	"survey_backend/models/serialization"
	"survey_backend/service"
	"survey_backend/utils"
)

func (LoginApi) Register(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.RegisterSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	if requestBody.Password != requestBody.ConfirmPassword {
		res.FailWithMsg("两次密码输入不一致！", c)
		return
	}

	// 验证用户是否重复
	_, ok := service.GetUserByParams(db.(*gorm.DB), requestBody.Phone)
	if ok {
		global.Log.Warn("用户已存在，请直接登录！")
		res.FailWithMsg("用户已存在，请直接登录！", c)
		return
	}

	userId, err := service.CreateUser(db.(*gorm.DB), &requestBody)
	if err != nil {
		res.FailWithMsg("注册失败", c)
		return
	}
	data := map[string]uint{"user_id": userId}
	res.OkWithData(data, c)
}

func (LoginApi) Login(c *gin.Context) {
	db, _ := c.Get("db")
	var requestBody serialization.LoginSerialization
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	// 验证用户是否存在
	user, ok := service.GetUserByParams(db.(*gorm.DB), requestBody.Username)
	if !ok {
		global.Log.Warn("用户名/手机号/邮箱不存在")
		res.FailWithCode(res.UserPwdError, c)
		return
	}
	// 校验密码
	if !utils.CheckPwd(requestBody.Password, user.Password) {
		global.Log.Warn("密码错误")
		res.FailWithCode(res.UserPwdError, c)
		return
	}
	// 设置token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		res.FailWithMsg("token生成失败", c)
		return
	}
	data := map[string]string{"token": token}
	res.OkWithData(data, c)
}
