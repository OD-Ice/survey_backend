package settings_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (SettingApi) SettingsInfoView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "abc"})
}
