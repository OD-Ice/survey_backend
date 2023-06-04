package logout_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"survey_backend/global"
	"survey_backend/models/res"
	"time"
)

func (LogoutApi) LogOut(c *gin.Context) {
	exp, _ := c.Get("exp")
	expTime := time.Unix(exp.(int64), 0)
	now := time.Now()

	diff := expTime.Sub(now)

	token := c.Request.Header.Get("token")

	err := global.Redis.Set(fmt.Sprintf("logout_%s", token), 1, diff).Err()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("注销失败", c)
		return
	}
	res.OkWith(c)
}
