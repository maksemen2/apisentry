package repository

import (
	"context"
	"github.com/maksemen2/apisentry/internal/core/entity"
)

type StatsRepository interface {
	GetStatsOverview(ctx context.Context) (*entity.OverviewStats, error)
	GetStatsByEndpoint(ctx context.Context, endpointID uint) (*entity.EndpointStats, error)
}
