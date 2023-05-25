package api

import "survey_backend/api/settings_api"

type ApiGroup struct {
	SettingApi settings_api.SettingApi
}

var ApiGroupApp = ApiGroup{}
