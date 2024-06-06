package utils

import "log"

type Alert struct {
    // Add fields as needed, such as email settings or webhook URLs
}

func NewAlert() *Alert {
    return &Alert{}
}

func (a *Alert) Send(message string) {
    // Placeholder for alerting logic, such as sending an email or a webhook
    log.Printf("ALERT: %s", message)
}
