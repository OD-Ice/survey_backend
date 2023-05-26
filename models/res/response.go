package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success = 0
	fail    = -1
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(success, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(success, data, "成功", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(success, map[string]any{}, msg, c)
}

func OkWith(c *gin.Context) {
	Result(success, map[string]any{}, "成功", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(fail, data, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(fail, map[string]any{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	Result(code.code, map[string]any{}, code.msg, c)
}
