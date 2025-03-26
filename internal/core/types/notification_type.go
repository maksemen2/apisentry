package types

type NotificationType string

const (
	EmailNotification    NotificationType = "email"
	WebhookNotification  NotificationType = "webhook"
	TelegramNotification NotificationType = "telegram"
)

func (n NotificationType) String() string {
	return string(n)
}

func (n NotificationType) IsValid() bool {
	switch n {
	case EmailNotification, WebhookNotification, TelegramNotification:
		return true
	}
	return false
}
