package backup

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmailWithURL(url string) error {
	to := os.Getenv("MAIL_TO")
	fromName := os.Getenv("MAIL_FROM_NAME")
	fromEmail := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	appName := os.Getenv("APP_NAME")

	// Message.
	subject := fmt.Sprintf("Subject:%s - Backup realizado com sucesso!\n", appName)
	body := fmt.Sprintf("Backup realizado com sucesso!\n\nArquivo: %s\n\nMRF Solution", url)

	message := []byte(fmt.Sprintf("From: %s <%s>\n%s\n%s", fromName, fromEmail, subject, body))

	// Authentication.
	auth := smtp.PlainAuth("", fromEmail, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, []string{to}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
