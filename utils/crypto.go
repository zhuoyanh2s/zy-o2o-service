package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5 加盐加密
func Md5(str string) string {
	h := md5.New()
	h.Write(append(secretKey, []byte(str)...))
	return hex.EncodeToString(h.Sum(nil))
}
