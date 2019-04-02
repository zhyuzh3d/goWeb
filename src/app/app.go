package main

import (
	"app/api"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	//获取当前程序运行的目录
	dir, _ := os.Getwd()
	webDir := path.Join(dir, "/web")

	//设置文件服务
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	//API-注册登录相关
	http.HandleFunc("/api/Register", api.Register)
	http.HandleFunc("/api/Login", api.Login)
	http.HandleFunc("/api/SendRegVerifyMail", api.SendRegVerifyMail)
	http.HandleFunc("/api/SendRstPwMail", api.SendRstPwMail)
	http.HandleFunc("/api/ResetPw", api.ResetPw)

	//启动服务
	fmt.Println("Server is running.Current Dictionary is", dir)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
