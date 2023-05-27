package enum

// 问题的类型
const (
	Single     = iota + 1 // 单选
	Multiple              // 多选
	Subjective            // 主观题
	Blanks                // 填空题
)
