package repositoryImplementation

import (
	"github.com/velo-project/events-service/adapters/dao/repositories"
	"github.com/velo-project/events-service/adapters/dao/tables/mappers"
	"github.com/velo-project/events-service/core/entities"
	"gorm.io/gorm"
)

type EventRepositoryImplementation struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) repositories.EventRepository {
	return &EventRepositoryImplementation{
		db: db,
	}
}

func (eri *EventRepositoryImplementation) CreateEvent(event entities.EventEntity) entities.EventEntity {
	persistanceEntity := mappers.EventToPersistanceTable(&event)

	eri.db.Create(persistanceEntity)

	return mappers.EventToDomainEntity(&persistanceEntity)
}
