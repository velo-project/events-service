package mappers

import (
	"database/sql"

	"github.com/velo-project/events-service/adapters/dao/tables"
	"github.com/velo-project/events-service/core/entities"
	"github.com/velo-project/events-service/core/entities/enums"
)

func EventToDomainEntity(eventTable tables.EventTable) entities.EventEntity {
	convertIntToBool := func(value int) bool {
		return value == 1
	}

	return entities.EventEntity{
		ID:                   int(eventTable.ID),
		Name:                 eventTable.Name,
		Description:          eventTable.Description.String,
		PhotoPath:            eventTable.PhotoPath.String,
		IsActive:             convertIntToBool(eventTable.IsActive),
		IsCanceled:           convertIntToBool(eventTable.IsCanceled),
		Address:              eventTable.Address,
		DifficultyOfTheTrack: enums.DifficultyEnum(eventTable.DifficultyOfTheTrack),
		TrackType:            enums.TrackTypeEnum(eventTable.TrackType),
		WhenWillHappen:       eventTable.WhenWillHappen,
		CreatedAt:            eventTable.CreatedAt,
		UpdatedAt:            eventTable.UpdatedAt,
	}
}

func EventToPersistanceTable(eventEntity entities.EventEntity) tables.EventTable {
	convertBoolToInt := func(value bool) int {
		if value {
			return 1
		}
		return 0
	}

	return tables.EventTable{
		ID:                   uint(eventEntity.ID),
		Name:                 eventEntity.Name,
		Description:          sql.NullString{String: eventEntity.Description, Valid: eventEntity.Description != ""},
		PhotoPath:            sql.NullString{String: eventEntity.PhotoPath, Valid: eventEntity.PhotoPath != ""},
		IsActive:             convertBoolToInt(eventEntity.IsActive),
		IsCanceled:           convertBoolToInt(eventEntity.IsCanceled),
		Address:              eventEntity.Address,
		DifficultyOfTheTrack: int(eventEntity.DifficultyOfTheTrack),
		TrackType:            int(eventEntity.TrackType),
		WhenWillHappen:       eventEntity.WhenWillHappen,
		CreatedAt:            eventEntity.CreatedAt,
		UpdatedAt:            eventEntity.UpdatedAt,
	}
}
