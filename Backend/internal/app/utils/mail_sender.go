package utils

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

var mail *gomail.Dialer

func ConfigMail() {
	d := gomail.NewDialer("smtp.gmail.com", 587, "hung18072002ht@gmail.com", "lrobejftgpgmjqez")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	mail = d
}

func SendMailOTP(to string, msg string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "hung18072002ht@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Mã OTP")
	m.SetBody("text/html", msg)
	err := mail.DialAndSend(m)
	if err != nil {
		fmt.Println("error sending mail: ", err)
	}
}
