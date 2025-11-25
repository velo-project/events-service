package services

import (
	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type getSubscribedEventsService struct {
	repository ports.GetSubscribedEventsRepository
}

func NewGetSubscribedEventsService(repository ports.GetSubscribedEventsRepository) ports.GetSubscribedEventsService {
	return &getSubscribedEventsService{
		repository: repository,
	}
}

func (s *getSubscribedEventsService) GetSubscribedEvents(userID string) ([]entities.Event, error) {
	return s.repository.GetSubscribedEvents(userID)
}
