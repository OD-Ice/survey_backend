package settings_api

import (
	"github.com/gin-gonic/gin"
	"survey_backend/models/res"
)

func (SettingApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(res.SystemError, c)
}
