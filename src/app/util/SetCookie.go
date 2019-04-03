package util

import (
	"net/http"
	"time"
)

//SetCookie 设置Cookie，默认1月/路径
func SetCookie(w http.ResponseWriter, k string, v string) {
	exp := time.Now().AddDate(0, 1, 0)
	path := "/"
	SetCookieExt(w, k, v, exp, path, 0, true)
}

//DelCookie 删除Cookie，MaxAge=-1
func DelCookie(w http.ResponseWriter, k string) {
	exp := time.Unix(0, 0)
	path := "/"
	SetCookieExt(w, k, "", exp, path, -1000, true)
}

//SetCookieExt 设置Cookie扩展版
func SetCookieExt(w http.ResponseWriter, k string, v string,
	exp time.Time, path string, max int, htp bool) {
	c := http.Cookie{
		Name:     k,
		Path:     path,
		Value:    v,
		HttpOnly: htp,
		Expires:  exp,
		MaxAge:   max,
	}
	http.SetCookie(w, &c)
}
