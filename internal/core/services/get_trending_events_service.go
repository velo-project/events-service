package services

import (
	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type getTrendingEventsService struct {
	repository ports.GetTrendingEventsRepository
}

func NewGetTrendingEventsService(repository ports.GetTrendingEventsRepository) ports.GetTrendingEventsService {
	return &getTrendingEventsService{
		repository: repository,
	}
}

func (s *getTrendingEventsService) GetTrendingEvents() ([]entities.Event, error) {
	return s.repository.GetTrendingEvents()
}
