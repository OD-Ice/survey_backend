package routers

import (
	"github.com/gin-gonic/gin"
	"survey_backend/global"
)

type RouterGroup struct {
	*gin.Engine
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	RouterGroup{router}.SettingsRouter()
	return router
}
