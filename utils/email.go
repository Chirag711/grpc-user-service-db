package utils

import (
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string) error {

	// fmt.Println("Email sent from ", os.Getenv("EMAIL_PASS"))

	m := gomail.NewMessage()

	m.SetHeader("From", "onlinedataacc@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Welcome to Our Platform")
	m.SetBody("text/html", `
		<h2>User Created Successfully</h2>
		<p>Thank you for registering.</p>
	`)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"onlinedataacc@gmail.com",
		"dsxqaqhbqyoevxzu",
	)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Email send error:", err)
		return err
	}

	return nil
}
