package services

import (
	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type getLastParticipatedEventsService struct {
	repository ports.GetLastParticipatedEventsRepository
}

func NewGetLastParticipatedEventsService(repository ports.GetLastParticipatedEventsRepository) ports.GetLastParticipatedEventsService {
	return &getLastParticipatedEventsService{
		repository: repository,
	}
}

func (s *getLastParticipatedEventsService) GetLastParticipatedEvents(userID string) ([]entities.Event, error) {
	return s.repository.GetLastParticipatedEvents(userID)
}
