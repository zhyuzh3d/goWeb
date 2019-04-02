package tool

import (
	"errors"
	"io"
	"io/ioutil"
	"log"

	"github.com/emersion/go-smtp"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	if username != "zhyuzh@myweb.com" || password != "password" {
		return nil, errors.New("Invalid username or password")
	}
	return &Session{}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

// A Session is returned after successful login.
type Session struct{}

//Mail ..
func (s *Session) Mail(from string) error {
	log.Println("Mail from:", from)
	return nil
}

//Rcpt ..
func (s *Session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	return nil
}

//Data ..
func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("Data:", string(b))
	}
	return nil
}

//Reset ..
func (s *Session) Reset() {}

//Logout ..
func (s *Session) Logout() error {
	return nil
}

func init() {
	go startServer()
}

func startServer() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":1025"
	s.Domain = "localhost"
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting mail server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
