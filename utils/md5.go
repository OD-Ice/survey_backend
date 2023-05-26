package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptMD5 使用 MD5 算法加密字符串，并返回加密后的结果
func EncryptMD5(message string) string {
	hash := md5.New()
	hash.Write([]byte(message))
	encrypted := hash.Sum(nil)
	return hex.EncodeToString(encrypted)
}

// VerifyMD5 校验原始字符串与加密后的字符串是否匹配
func VerifyMD5(original, encrypted string) bool {
	return EncryptMD5(original) == encrypted
}
