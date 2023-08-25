package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	// Get the MAILTRAP_USERNAME environment variable
	mailtrapUsername, exists := os.LookupEnv("MAILTRAP_USERNAME")
	fmt.Println(mailtrapUsername)
	if !exists {
		fmt.Println("no username")
	}

	// Get the GITHUB_API_KEY environment variable
	mailtrapPassword, exists := os.LookupEnv("MAILTRAP_PASSWORD")

	if !exists {
		fmt.Println("no password")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	s := io.WriterTo(m)

	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, mailtrapUsername, mailtrapPassword)
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
