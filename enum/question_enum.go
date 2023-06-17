package enum

// 问题的类型
const (
	Single     = iota + 1 // 单选
	Multiple              // 多选
	Subjective            // 主观题
)

// 问卷状态
const ( // 正常
	Published = 1 // 发布
)
