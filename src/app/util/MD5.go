package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

//MD5 加密字符串
func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	a := h.Sum(nil)
	s := hex.EncodeToString(a)
	return s
}
