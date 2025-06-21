package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/mail.v2"
)

func SendOTPEmail(email, otpCode string) error {
	// In development, just log the OTP
	if os.Getenv("APP_ENV") == "development" {
		fmt.Printf("📧 OTP for %s: %s\n", email, otpCode)
		return nil
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	if smtpHost == "" || smtpUser == "" || smtpPass == "" {
		fmt.Printf("📧 Email not configured, OTP for %s: %s\n", email, otpCode)
		return nil
	}

	smtpPort := 587
	if smtpPortStr != "" {
		if port, err := strconv.Atoi(smtpPortStr); err == nil {
			smtpPort = port
		}
	}

	m := mail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "VCM Medical Platform - Verification Code")
	
	body := fmt.Sprintf(`
		<h2>VCM Medical Platform</h2>
		<p>Your verification code is: <strong>%s</strong></p>
		<p>This code will expire in 10 minutes.</p>
		<p>If you didn't request this code, please ignore this email.</p>
	`, otpCode)
	
	m.SetBody("text/html", body)

	d := mail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	
	return d.DialAndSend(m)
}
