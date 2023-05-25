package routers

import (
	"github.com/gin-gonic/gin"
	"survey_backend/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "测试")
	})
	return router
}
