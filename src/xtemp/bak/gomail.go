package main

import "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "from@example.com")
	m.SetHeader("To", "ylzhang@gem-inno.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello!")

	d := gomail.Dialer{Host: "localhost", Port: 587}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
