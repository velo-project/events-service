package services

import (
	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type getRecommendedEventsService struct {
	repository ports.GetRecommendedEventsRepository
}

func NewGetRecommendedEventsService(repository ports.GetRecommendedEventsRepository) ports.GetRecommendedEventsService {
	return &getRecommendedEventsService{
		repository: repository,
	}
}

func (s *getRecommendedEventsService) GetRecommendedEvents(userID string) ([]entities.Event, error) {
	return s.repository.GetRecommendedEvents(userID)
}
