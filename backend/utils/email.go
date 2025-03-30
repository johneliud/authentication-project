package utils

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

// SendVerificationEmail sends an email to the user with a verification code.
func SendVerificationEmail(userEmail, verificationCode string) error {
	e := email.NewEmail()
	e.From = "Authentication Project <johneliud4@gmail.com>"
	e.To = []string{userEmail}
	e.Subject = "Email Verification Code"
	e.Text = []byte(fmt.Sprintf("Your verification code is: %s", verificationCode))

	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), smtpServer)

	if err := e.Send(smtpServer+":"+smtpPort, auth); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
