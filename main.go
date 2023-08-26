package main

import (
	"fmt"
	"gogo/config"
	"io"
	"log"

	"gopkg.in/gomail.v2"
)

var MailtrapUser string
var MailtrapPassword string

func init() {
	config.SetEnv()
	MailtrapUser = config.ApiConfig.Mailtraps[0].User
	MailtrapPassword = config.ApiConfig.Mailtraps[0].Password
}

func main() {
	fmt.Println(MailtrapUser, MailtrapPassword)
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	s := io.WriterTo(m)
	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, MailtrapUser, MailtrapPassword)

	conn, err := d.Dial()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	to := []string{"alex@example.com"}

	// Send the email and capture the response
	if err := conn.Send("alex@example.com", to, s); err != nil {
		fmt.Println("Email err", err)
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully!")

}
