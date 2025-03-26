package repository

import (
	"context"

	"github.com/maksemen2/apisentry/internal/core/entity"
)

// EndpointRepository is an interface for managing endpoints
type EndpointRepository interface {
	Update(ctx context.Context, endpoint *entity.Endpoint) (*entity.Endpoint, error)
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*entity.Endpoint, error)
	GetAll(ctx context.Context)
	Create(ctx context.Context, endpoint *entity.Endpoint) (*entity.Endpoint, error)
}
