package mappers

import (
	"database/sql"

	"github.com/velo-project/events-service/adapters/dao/tables"
	"github.com/velo-project/events-service/core/entities"
)

func EventToDomainEntity(eventTable *tables.EventTable) entities.EventEntity {
	return entities.EventEntity{
		ID:          int(eventTable.ID),
		Name:        eventTable.Name,
		Description: eventTable.Description.String,
		PhotoPath:   eventTable.PhotoPath.String,
	}
}

func EventToPersistanceTable(eventEntity *entities.EventEntity) tables.EventTable {
	return tables.EventTable{
		ID:          uint(eventEntity.ID),
		Name:        eventEntity.Name,
		Description: sql.NullString{String: eventEntity.Description, Valid: eventEntity.Description != ""},
		PhotoPath:   sql.NullString{String: eventEntity.PhotoPath, Valid: eventEntity.PhotoPath != ""},
	}
}
