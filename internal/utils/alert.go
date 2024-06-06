package utils

import (
    "gopkg.in/mail.v2"
    "log"
)

// AlertConfig contains configuration for sending alerts
type AlertConfig struct {
    SMTPHost     string
    SMTPPort     int
    SMTPUsername string
    SMTPPassword string
    Recipient    string
}

// SendEmail sends an email alert
func SendEmail(cfg AlertConfig, subject, body string) error {
    m := mail.NewMessage()
    m.SetHeader("From", cfg.SMTPUsername)
    m.SetHeader("To", cfg.Recipient)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := mail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUsername, cfg.SMTPPassword)

    if err := d.DialAndSend(m); err != nil {
        return err
    }

    return nil
}

func SendAlert(subject, message string) {
    cfg := AlertConfig{
        SMTPHost:     "smtp.example.com",  // Replace with your SMTP server
        SMTPPort:     587,
        SMTPUsername: "your_email@example.com",  // Replace with your email
        SMTPPassword: "your_password",           // Replace with your email password
        Recipient:    "a9711880352@gmail.com",
    }

    err := SendEmail(cfg, subject, message)
    if err != nil {
        log.Printf("Failed to send alert: %v", err)
    } else {
        log.Println("Alert sent successfully")
    }
}
