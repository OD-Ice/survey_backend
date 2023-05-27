package api

import (
	"survey_backend/api/questionnaire_api"
	"survey_backend/api/settings_api"
)

type ApiGroup struct {
	SettingApi       settings_api.SettingApi
	QuestionnaireApi questionnaire_api.QuestionnaireApi
	QuestionApi      questionnaire_api.QuestionApi
}

var ApiGroupApp = ApiGroup{}
