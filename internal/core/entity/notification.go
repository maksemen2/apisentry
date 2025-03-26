package entity

import (
	coreerrors "github.com/maksemen2/apisentry/internal/core/errors"
	"github.com/maksemen2/apisentry/internal/core/types"
)

type Notification struct {
	ID       uint
	Type     types.NotificationType
	Email    string
	ChatID   int64
	BotToken string
	URL      string
}

func (e *Notification) Validate() error {

	var errors []error

	if !e.Type.IsValid() {
		errors = append(errors, coreerrors.ErrInvalidNotificationType)
	}
	if e.Type == types.EmailNotification && e.Email == "" {
		errors = append(errors, coreerrors.ErrEmailRequired)
	}
	if e.Type == types.TelegramNotification {
		if e.ChatID == 0 {
			errors = append(errors, coreerrors.ErrChatIDRequired)
		}
		if e.BotToken == "" {
			errors = append(errors, coreerrors.ErrBotTokenRequired)
		}
	}
	if e.Type == types.WebhookNotification && e.URL == "" {
		errors = append(errors, coreerrors.ErrWebhookURLRequired)
	}

	if len(errors) > 0 {
		return coreerrors.NewValidationErrors(errors)
	}

	return nil
}
