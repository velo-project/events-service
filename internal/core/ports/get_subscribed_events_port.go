package ports

import "gitlab.com/velo-company/services/events-service/internal/core/entities"

type GetSubscribedEventsRepository interface {
	GetSubscribedEvents(userID string) ([]entities.Event, error)
}

type GetSubscribedEventsService interface {
	GetSubscribedEvents(userID string) ([]entities.Event, error)
}
