package main

import (
	"os"

	"github.com/bashmc/draft/mail"
)

type Config struct {
	MailConfig  *mail.Config
	PostgresURL string
}

func loadConfig() *Config {

	mailCfg := &mail.Config{
		Host:        os.Getenv("MAIL_HOST_URL"),
		Token:       os.Getenv("MAIL_SECRET_TOKEN"),
		SenderEmail: os.Getenv("MAIL_SENDER_EMAIL"),
		SenderName:  os.Getenv("MAIL_SENDER_NAME"),
	}

	return &Config{
		MailConfig:  mailCfg,
		PostgresURL: os.Getenv("DB_URL"),
	}
}