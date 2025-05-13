package emai

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/smtp"
	"os"
)

func SendEmail(email, name string) error {

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	from := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", from, pass, host)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: We are very grateful\r\n"+
			"MIME-version: 1.0\r\n"+
			"Content-Type: text/plain; charset=\"UTF-8\"\r\n"+
			"\r\n"+
			"Our team is very grateful for your donation, %s",
		from, email, name))

	addr := host + ":" + port
	return smtp.SendMail(addr, auth, from, []string{email}, msg)
}
