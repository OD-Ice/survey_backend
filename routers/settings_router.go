package routers

import (
	"survey_backend/api"
)

func (router RouterGroup) SettingsRouter() {
	// 分组
	settingsRouter := router.Group("settings")
	settingsApi := api.ApiGroupApp.SettingApi
	settingsRouter.GET("/", settingsApi.SettingsInfoView)
}
