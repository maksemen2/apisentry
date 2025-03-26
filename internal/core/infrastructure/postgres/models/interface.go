package postgresmodels

import (
	"github.com/maksemen2/apisentry/internal/core/entity"
	"gorm.io/gorm"
)

type postgresModel interface {
	gorm.Model
	ToEntity() (entity.Entity, error)
}
