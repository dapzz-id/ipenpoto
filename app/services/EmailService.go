package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"ipenpoto/database/models"
)

type EmailService struct {
	SmtpHost   string
	SmtpPort   int
	SmtpUser   string
	SmtpPass   string
	SenderName string
}

func NewEmailService() *EmailService {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if port == 0 {
		port = 587
	}

	return &EmailService{
		SmtpHost:   os.Getenv("SMTP_HOST"),
		SmtpPort:   port,
		SmtpUser:   os.Getenv("SMTP_USER"),
		SmtpPass:   os.Getenv("SMTP_PASS"),
		SenderName: os.Getenv("SMTP_SENDER_NAME"),
	}
}

func (s *EmailService) SendVerificationEmail(
	user *models.User,
	verificationToken string,
) error {
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	verificationLink := fmt.Sprintf(
		"%s/verify-email?token=%s",
		frontendURL,
		verificationToken,
	)

	isMockMode := s.SmtpUser == "" || s.SmtpPass == "" ||
		s.SmtpUser == "isi_dengan_username_anda" ||
		s.SmtpPass == "isi_dengan_password_anda"

	if isMockMode {
		log.Printf("[MOCK EMAIL] Verification email for %s (%s)\n", user.Email, user.Name)
		log.Printf("[MOCK EMAIL] Verification link: %s\n", verificationLink)
		log.Printf("[MOCK EMAIL] Token: %s\n", verificationToken)
		return nil
	}

	subject := "Verify Your Email - Ipenpoto"
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background-color: #007bff; color: white; padding: 20px; text-align: center; border-radius: 5px 5px 0 0; }
        .content { background-color: #f9f9f9; padding: 20px; border: 1px solid #ddd; border-radius: 0 0 5px 5px; }
        .button { display: inline-block; padding: 12px 30px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px; margin: 20px 0; }
        .footer { margin-top: 20px; font-size: 12px; color: #666; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Welcome to Ipenpoto!</h1>
        </div>
        <div class="content">
            <p>Hi <strong>%s</strong>,</p>
            <p>Thank you for registering with Ipenpoto. Please verify your email address to activate your account.</p>
            <p>Click the button below to verify your email:</p>
            <a href="%s" class="button">Verify Email</a>
            <p>Or copy and paste this link in your browser:</p>
            <p style="word-break: break-all; font-size: 12px;">%s</p>
            <p>This verification link will expire in 24 hours.</p>
            <p>If you didn't create this account, please ignore this email.</p>
            <div class="footer">
                <p>Best regards,<br>Ipenpoto Team</p>
            </div>
        </div>
    </div>
</body>
</html>
`,
		user.Name,
		verificationLink,
		verificationLink,
	)

	from := fmt.Sprintf("%s <%s>", s.SenderName, s.SmtpUser)
	to := []string{user.Email}

	headers := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n",
		from,
		user.Email,
		subject,
	)

	message := headers + "\r\n" + body

	auth := smtp.PlainAuth(
		"",
		s.SmtpUser,
		s.SmtpPass,
		s.SmtpHost,
	)

	addr := fmt.Sprintf("%s:%d", s.SmtpHost, s.SmtpPort)

	err := smtp.SendMail(
		addr,
		auth,
		s.SmtpUser,
		to,
		[]byte(message),
	)

	return err
}
