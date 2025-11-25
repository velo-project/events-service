package ports

import "gitlab.com/velo-company/services/events-service/internal/core/entities"

type GetTrendingEventsRepository interface {
	GetTrendingEvents() ([]entities.Event, error)
}

type GetTrendingEventsService interface {
	GetTrendingEvents() ([]entities.Event, error)
}
