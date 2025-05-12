package emai

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(email, name string) error {
	from := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", from, pass, host)

	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: We are very grateful \r\n"+
			"\r\n"+
			"Our big team very greatful for you donation %s",
		email, name))

	addr := host + ":" + port
	return smtp.SendMail(addr, auth, from, []string{email}, msg)
}
