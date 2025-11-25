package ports

import "gitlab.com/velo-company/services/events-service/internal/core/entities"

type GetLastParticipatedEventsRepository interface {
	GetLastParticipatedEvents(userID string) ([]entities.Event, error)
}

type GetLastParticipatedEventsService interface {
	GetLastParticipatedEvents(userID string) ([]entities.Event, error)
}
