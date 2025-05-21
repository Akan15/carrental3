package email

import "fmt"

func SendEmail(to, subject, body string) error {
	fmt.Println("📧 MOCK EMAIL")
	fmt.Println("To:", to)
	fmt.Println("Subject:", subject)
	fmt.Println("Body:", body)
	return nil
}

//package email

//import (
//	"net/smtp"
//	"os"
//)

//func SendEmail(to, subject, body string) error {
//	from := os.Getenv("SMTP_FROM")
//	password := os.Getenv("SMTP_PASS")

//	smtpHost := "smtp.mail.ru"
//	smtpPort := "587"

//	auth := smtp.PlainAuth("", from, password, smtpHost)

//	msg := []byte("To: " + to + "\r\n" +
//		"Subject: " + subject + "\r\n\r\n" +
//		body + "\r\n")

//	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
//}
