package repository

import (
	"context"
	"github.com/maksemen2/apisentry/internal/core/entity"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *entity.Notification) (*entity.Notification, error)
	GetByEndpointID(ctx context.Context, endpointID uint) ([]entity.Notification, error)
}
