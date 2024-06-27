package utils

import (
	"crypto/tls"
	"fmt"
	"music-app-backend/internal/app/config"

	"gopkg.in/gomail.v2"
)

type MailSender struct {
	mail *gomail.Dialer
}

func NewMailSender(config *config.Config) *MailSender {
	return &MailSender{
		mail: ConfigMail(config),
	}
}

func ConfigMail(config *config.Config) *gomail.Dialer {
	d := gomail.NewDialer(
		config.MailSenderHost,
		config.MailSenderPort,
		config.MailSenderUsername,
		config.MailSenderPassword,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d
}

func (mail *MailSender) SendMailOTP(to string, msg string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "hung18072002ht@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Mã OTP")
	m.SetBody("text/html", msg)
	err := mail.mail.DialAndSend(m)
	if err != nil {
		fmt.Println("error sending mail: ", err)
	}
}

func (mail *MailSender) SendMail(to string, msg string, subject string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "hung18072002ht@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)
	err := mail.mail.DialAndSend(m)
	if err != nil {
		fmt.Println("error sending mail: ", err)
	}
}
