package main

import (
	"app/api"
	"app/ext"
	"fmt"
	"net/http"
	"os"
	"path"
)

func main() {
	//获取当前程序运行的目录
	dir, _ := os.Getwd()
	webDir := path.Join(dir, "/web")

	//文件服务器和中间件
	fileHandler := http.FileServer(http.Dir(webDir))
	http.Handle("/", ext.MiddleWare(fileHandler))

	//API-注册登录相关
	http.HandleFunc("/api/Register", api.Register)
	http.HandleFunc("/api/Login", api.Login)
	http.HandleFunc("/api/SendRegVerifyMail", api.SendRegVerifyMail)
	http.HandleFunc("/api/SendRstPwMail", api.SendRstPwMail)
	http.HandleFunc("/api/ResetPw", api.ResetPw)
	http.HandleFunc("/api/AutoLogin", ext.MiddleWareAPI(api.AutoLogin))

	//启动服务
	fmt.Println("Server is running.Current Dictionary is", dir)
	http.ListenAndServe(":8088", nil)
}
