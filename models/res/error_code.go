package res

type ErrorCode struct {
	code int
	msg  string
}

var (
	SystemError    = ErrorCode{1001, "系统错误"}
	ParameterError = ErrorCode{1002, "参数错误"}
)
