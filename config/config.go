package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Mailtrap struct {
	Password string
	User     string
}

type Config struct {
	Mailtraps []Mailtrap
	Logs      []Log
}

type Log struct {
	Type string
	Path string
}

var ApiConfig Config

func SetEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Environment error")
	}

	MailtrapPassword := os.Getenv("MAILTRAP_PASSWORD")
	MailtrapUser := os.Getenv("MAILTRAP_USER")
	firstMailtrap := Mailtrap{Password: MailtrapPassword, User: MailtrapUser}

	FirstLog := Log{"Fatal", "var/www/logs/fatal.txt"}
	SecondLog := Log{"Error", "var/www/data/errors/error.txt"}

	ApiConfig = Config{
		Mailtraps: []Mailtrap{firstMailtrap},
		Logs:      []Log{FirstLog, SecondLog},
	}

	thirdLog := Log{"Info", "var/www/info/info.txt"}

	ApiConfig.Logs = append(ApiConfig.Logs, thirdLog)
}
