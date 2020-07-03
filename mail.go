package godash

import (
	"log"
	"gopkg.in/gomail.v2"
)

const mailFrom = "example@example.com"

func SendMail(email string, title string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "Example"+"<"+mailFrom+">")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", title)
	msg.SetBody("text/html", body)

	//
	mailer := gomail.NewDialer(
		"smtp.example.com",
		465,
		"example@example.com",
		"example",
	)
	if err := mailer.DialAndSend(msg); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
