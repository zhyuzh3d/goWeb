package main

import (
	"log"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// return "LOGIN", []byte{}, nil
	return "LOGIN", []byte(a.username), nil
}
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

func main() {
	// auth := smtp.PlainAuth("", "zhyuzhxd@hotmail.com", "zhyuzh3d", "smtp.office365.com")
	auth := LoginAuth("zhyuzhxd@hotmail.com", "zhyuzh3d")
	to := []string{"ylzhang@gem-inno.com"}
	/* msg := []byte("To: ylzhang@gem-inno.com\r\n" +
	"Subject: discount Gophers!\r\n" +
	"\r\n" +
	"This is the email body.\r\n") */

	content_type := "Content-Type: text/plain" + "; charset=UTF-8"
	msg := []byte("To: ylzhang@gem-inno.com" +
		"\r\nFrom: zhyuzhxd@hotmail.com" +
		"\r\nSubject: Hello" +
		"\r\n" + content_type + "\r\n\r\n" +
		"Hello world!")
	err := smtp.SendMail("smtp.office365.com:587", auth, "zhyuzhxd@hotmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
