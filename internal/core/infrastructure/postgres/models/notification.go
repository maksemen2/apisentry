package postgresmodels

import (
	"github.com/maksemen2/apisentry/internal/core/entity"
	coreerrors "github.com/maksemen2/apisentry/internal/core/errors"
	"github.com/maksemen2/apisentry/internal/core/types"
	"gorm.io/gorm"
)

type notificationModel struct {
	gorm.Model
	Type     string
	Email    string
	ChatID   int64
	BotToken string
	URL      string
}

func (n *notificationModel) ToEntity() (*entity.Notification, error) {

	notificationType := types.NotificationType(n.Type)
	if !notificationType.IsValid() {
		return nil, coreerrors.ErrInvalidNotificationType
	}

	return &entity.Notification{
		ID:       n.ID,
		Type:     notificationType,
		Email:    n.Email,
		ChatID:   n.ChatID,
		BotToken: n.BotToken,
		URL:      n.URL,
	}, nil
}
