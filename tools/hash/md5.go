package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5 快速计算Md5
func Md5(str string) string {
	hashInstance := md5.New()
	hashInstance.Write([]byte(str))
	return hex.EncodeToString(hashInstance.Sum(nil))
}

// Md5Check 校验Md5
func Md5Check(str string, md5 string) bool {
	return md5 == Md5(str)
}

// ByteMd5 传递[]byte计算md5
func ByteMd5(data []byte) string {
	hashInstance := md5.New()
	hashInstance.Write(data)
	return hex.EncodeToString(hashInstance.Sum(nil))
}

// FileMd5 计算文件Md5
func FileMd5(file *os.File) string {
	hashInstance := md5.New()
	io.Copy(hashInstance, file)
	return hex.EncodeToString(hashInstance.Sum(nil))
}
