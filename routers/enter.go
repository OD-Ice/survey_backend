package routers

import (
	"github.com/gin-gonic/gin"
	"survey_backend/global"
	"survey_backend/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(middleware.ReqHeaders())
	router.Use(middleware.ApiMySql())
	baseRouter := router.Group("api")
	RouterGroup{baseRouter}.SettingsRouter()
	RouterGroup{baseRouter}.QuestionnaireRouter()
	RouterGroup{baseRouter}.QuestionRouter()
	RouterGroup{baseRouter}.LoginRouter()
	RouterGroup{baseRouter}.LogoutRouter()
	return router
}
