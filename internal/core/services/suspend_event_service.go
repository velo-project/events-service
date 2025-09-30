package services

import (
	goErrors "errors"
	"gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type SuspendEventService interface {
	Execute(input *SuspendEventServiceInput) *SuspendEventServiceOutput
}

type suspendEventService struct {
	suspendEventPort   ports.SuspendEventPort
	userExistsByIdPort ports.UserExistsByIdPort
}

func NewSuspendEventService(suspendEventPort ports.SuspendEventPort, userExistsByIdPort ports.UserExistsByIdPort) SuspendEventService {
	return &suspendEventService{suspendEventPort: suspendEventPort, userExistsByIdPort: userExistsByIdPort}
}

type SuspendEventServiceInput struct {
	EventId int
	UserId  int
}

type SuspendEventServiceOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (s suspendEventService) Execute(input *SuspendEventServiceInput) *SuspendEventServiceOutput {
	exists, err := s.userExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return &SuspendEventServiceOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return &SuspendEventServiceOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	err = s.suspendEventPort.Execute(input.EventId)
	if err != nil {
		if goErrors.Is(err, errors.ErrEventNotFound) {
			return &SuspendEventServiceOutput{
				Message:    errors.ErrEventNotFound.Error(),
				StatusCode: 404,
			}
		}

		if goErrors.Is(err, errors.ErrBlockedSuspendEvent) {
			return &SuspendEventServiceOutput{
				Message:    errors.ErrBlockedSuspendEvent.Error(),
				StatusCode: 400,
			}
		}

		return &SuspendEventServiceOutput{
			Message:    "Não foi possível suspender o evento",
			StatusCode: 500,
		}
	}

	return &SuspendEventServiceOutput{
		Message:    "OK",
		StatusCode: 200,
	}
}
