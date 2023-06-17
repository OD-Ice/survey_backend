package utils

// InList 判断元素是否在列表中
func InList(key any, list []any) bool {
	for _, item := range list {
		if key == item {
			return true
		}
	}
	return false
}

// IsUintListEmpty 判断uint列表是否是空的
func IsUintListEmpty(slice []uint) bool {
	return len(slice) == 0
}
