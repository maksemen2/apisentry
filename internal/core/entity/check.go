package entity

import coreerrors "github.com/maksemen2/apisentry/internal/core/errors"

type Check struct {
	EndpointID   uint
	StatusCode   int
	ResponseTime float32
	Success      bool
	ErrorMessage string
}

func (e *Check) Validate() error {

	var errors []error

	if e.ResponseTime < 0 {
		errors = append(errors, coreerrors.ErrInvalidResponseTime)
	}

	if len(errors) > 0 {
		return coreerrors.NewValidationErrors(errors)
	}
	return nil
}
