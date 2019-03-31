package main

import (
	"app/api/login"
	"app/api/register"
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
	http.HandleFunc("/api/register", register.HandleFunc)
	http.HandleFunc("/api/login", login.HandleFunc)

	//启动服务
	fmt.Println("Server is running;Current Dir is", dir)
	log.Fatal(http.ListenAndServe(":8088", nil))

}
