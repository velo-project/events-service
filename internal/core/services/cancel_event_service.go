package services

import (
	"gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type CancelEventService interface {
	Execute(input *CancelEventServiceInput) *CancelEventServiceOutput
}

type cancelEventService struct {
	repo               ports.CancelEventPort
	userExistsByIdPort ports.UserExistsByIdPort
}

func NewCancelEventService(repo ports.CancelEventPort, userExistsByIdPort ports.UserExistsByIdPort) CancelEventService {
	return &cancelEventService{repo: repo, userExistsByIdPort: userExistsByIdPort}
}

type CancelEventServiceInput struct {
	EventId int
	UserId  int
}

type CancelEventServiceOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (s cancelEventService) Execute(input *CancelEventServiceInput) *CancelEventServiceOutput {
	exists, err := s.userExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return &CancelEventServiceOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return &CancelEventServiceOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	event, err := s.repo.GetEventById(input.EventId)
	if err != nil {
		return &CancelEventServiceOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			StatusCode: 500,
		}
	}

	if event == nil {
		return &CancelEventServiceOutput{
			Message:    errors.ErrEventNotFound.Error(),
			StatusCode: 404,
		}
	}

	err = s.repo.CancelEvent(input.EventId)
	if err != nil {
		return &CancelEventServiceOutput{
			Message:    "Não foi possível cancelar o evento",
			StatusCode: 500,
		}
	}

	return &CancelEventServiceOutput{
		Message:    "OK",
		StatusCode: 200,
	}
}
