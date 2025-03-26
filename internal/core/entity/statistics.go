package entity

import coreerrors "github.com/maksemen2/apisentry/internal/core/errors"

type OverviewStats struct {
	TotalEndpoints   int
	UptimePercentage float32
}

func (e *OverviewStats) Validate() error {

	errors := []error{}

	if e.TotalEndpoints < 0 {
		errors = append(errors, coreerrors.ErrInvalidTotalEndpoints)
	}
	if e.UptimePercentage < 0 {
		errors = append(errors, coreerrors.ErrInvalidUptimePercentage)
	}

	if len(errors) > 0 {
		return coreerrors.NewValidationErrors(errors)
	}

	return nil
}

type EndpointStats struct {
	Uptime24Hours   float32
	AvgResponseTime float32
}

func (e *EndpointStats) Validate() error {

	var errors []error

	if e.Uptime24Hours < 0 {
		errors = append(errors, coreerrors.ErrInvalidUptime24Hours)
	}
	if e.AvgResponseTime < 0 {
		errors = append(errors, coreerrors.ErrInvalidAvgResponseTime)
	}

	if len(errors) > 0 {
		return coreerrors.NewValidationErrors(errors)
	}

	return nil
}
