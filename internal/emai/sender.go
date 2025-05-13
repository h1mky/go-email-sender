package emai

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendEmail(email, name string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	from := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")

	var port int

	port, err = strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %v", err)
	}

	m := gomail.NewMessage()

	// Заголовки письма
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "We are very grateful")

	m.SetBody("text/html", fmt.Sprintf(
		"<p>Dear %s,</p><p>Our team is very grateful for your donation.</p>", name))

	d := gomail.NewDialer(host, port, from, pass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
