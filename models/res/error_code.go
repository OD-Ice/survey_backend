package res

type ErrorCode struct {
	code int
	msg  string
}

var (
	SystemError = ErrorCode{1001, "系统错误"}
)
