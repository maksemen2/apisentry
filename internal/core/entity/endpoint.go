package entity

import (
	coreerrors "github.com/maksemen2/apisentry/internal/core/errors"
	"github.com/maksemen2/apisentry/internal/core/types"
	"time"
)

type Endpoint struct {
	ID              uint
	URL             string
	IntervalSeconds time.Duration
	RetryStrategy   types.RetryStrategy
	MaxRetries      int
	Status          types.EndpointStatus
	CreatedAt       time.Time
}

func (e *Endpoint) Validate() error {

	var errors []error

	if e.IntervalSeconds <= 0 {
		errors = append(errors, coreerrors.ErrInvalidInterval)
	}
	if !e.RetryStrategy.IsValid() {
		errors = append(errors, coreerrors.ErrInvalidRetryStrategy)
	}
	if e.MaxRetries < 0 {
		errors = append(errors, coreerrors.ErrInvalidMaxRetries)
	}
	if !e.Status.IsValid() {
		errors = append(errors, coreerrors.ErrInvalidStatus)
	}

	if len(errors) > 0 {
		return coreerrors.NewValidationErrors(errors)
	}

	return nil
}
