package repositories

import "github.com/velo-project/events-service/core/entities"

type EventRepository interface {
	CreateEvent(event entities.EventEntity) entities.EventEntity
}
