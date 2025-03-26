package repository

import (
	"context"
	"github.com/maksemen2/apisentry/internal/core/entity"
)

// CheckRepository is an interface for managing checks
type CheckRepository interface {
	Create(ctx context.Context, check *entity.Check) (*entity.Check, error)
	GetAll(ctx context.Context, check *entity.Check) ([]entity.Check, error)
}
