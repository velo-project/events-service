package ports

import "gitlab.com/velo-company/services/events-service/internal/core/entities"

type GetRecommendedEventsRepository interface {
	GetRecommendedEvents(userID string) ([]entities.Event, error)
}

type GetRecommendedEventsService interface {
	GetRecommendedEvents(userID string) ([]entities.Event, error)
}
