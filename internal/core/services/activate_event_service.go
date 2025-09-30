package services

import (
	goErrors "errors"
	"gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type ActivateEventService interface {
	Execute(input *ActivateEventServiceInput) *ActivateEventServiceOutput
}

type activateEventService struct {
	activateEventPort  ports.ActivateEventPort
	userExistsByIdPort ports.UserExistsByIdPort
}

func NewActivateEventService(activateEventPort ports.ActivateEventPort, userExistsByIdPort ports.UserExistsByIdPort) ActivateEventService {
	return &activateEventService{activateEventPort: activateEventPort, userExistsByIdPort: userExistsByIdPort}
}

type ActivateEventServiceInput struct {
	EventId int
	UserId  int
}

type ActivateEventServiceOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (s activateEventService) Execute(input *ActivateEventServiceInput) *ActivateEventServiceOutput {
	exists, err := s.userExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return &ActivateEventServiceOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return &ActivateEventServiceOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	err = s.activateEventPort.Execute(input.EventId)
	if err != nil {
		if goErrors.Is(err, errors.ErrEventNotFound) {
			return &ActivateEventServiceOutput{
				Message:    errors.ErrEventNotFound.Error(),
				StatusCode: 404,
			}
		}

		if goErrors.Is(err, errors.ErrBlockedActivateEvent) {
			return &ActivateEventServiceOutput{
				Message:    errors.ErrBlockedActivateEvent.Error(),
				StatusCode: 400,
			}
		}

		return &ActivateEventServiceOutput{
			Message:    "Não foi possível ativar o evento",
			StatusCode: 500,
		}
	}

	return &ActivateEventServiceOutput{
		Message:    "OK",
		StatusCode: 200,
	}
}
