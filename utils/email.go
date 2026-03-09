package utils

import "gopkg.in/mail.v2"

func SendEmail(to string) error {

	m := mail.NewMessage()

	m.SetHeader("From", "your_email@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "User Created")
	m.SetBody("text/plain", "Your account has been created successfully")

	d := mail.NewDialer("smtp.gmail.com", 587, "your_email@gmail.com", "your_app_password")

	return d.DialAndSend(m)
}
