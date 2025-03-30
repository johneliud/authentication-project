package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

// SendVerificationEmail sends an email to the user with a verification code.
func SendVerificationEmail(userEmail, verificationCode string) error {
	e := email.NewEmail()

	e.From = "Authentication Project <" + os.Getenv("SMTP_EMAIL") + ">"
	e.To = []string{userEmail}
	e.Subject = "Email Verification Code"
	e.Text = []byte(fmt.Sprintf("Your verification code is: %s", verificationCode))

	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	auth := smtp.PlainAuth("", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), smtpServer)

	if err := e.Send(smtpServer+":"+smtpPort, auth); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}

// GenerateVerificationCode generates a random verification code.
func GenerateVerificationCode() string {
	bytes := make([]byte, 3)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
