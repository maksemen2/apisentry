package postgresmodels

import (
	"github.com/maksemen2/apisentry/internal/core/entity"
	"github.com/maksemen2/apisentry/internal/core/types"
	"gorm.io/gorm"
	"time"
)

type endpointModel struct {
	gorm.Model
	URL           string `gorm:"not null"`
	Interval      int
	RetryStrategy string `gorm:"type:varchar(20);not null"`
	Status        string `gorm:"index"`
	MaxRetries    int
}

func (e *endpointModel) ToEntity() (*entity.Endpoint, error) {
	endpoint := &entity.Endpoint{
		ID:              e.ID,
		URL:             e.URL,
		IntervalSeconds: time.Duration(e.Interval) * time.Second,
		RetryStrategy:   types.RetryStrategy(e.RetryStrategy),
		MaxRetries:      e.MaxRetries,
		Status:          types.EndpointStatus(e.Status),
		CreatedAt:       e.CreatedAt,
	}

	if err := endpoint.Validate(); err != nil {
		return nil, err
	}

	return endpoint, nil

}
