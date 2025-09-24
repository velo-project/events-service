package services

import (
	"errors"

	domainErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type SubscribeEventService interface {
	Execute(input *SubscribeEventServiceInput) *SubscribeEventServiceOutput
}

type subscribeEventService struct {
	UserExistsByIdPort ports.UserExistsByIdPort
	SubscribeEventPort ports.SubscribeEventPort
}

type SubscribeEventServiceInput struct {
	UserId  int
	EventId int
}

type SubscribeEventServiceOutput struct {
	Message    string  `json:"message"`
	Code       *string `json:"code"`
	StatusCode int     `json:"status_code"`
}

func NewSubscribeEventService(ue ports.UserExistsByIdPort, se ports.SubscribeEventPort) SubscribeEventService {
	return &subscribeEventService{
		UserExistsByIdPort: ue,
		SubscribeEventPort: se,
	}
}

func (s subscribeEventService) Execute(input *SubscribeEventServiceInput) *SubscribeEventServiceOutput {
	exists, err := s.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return &SubscribeEventServiceOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return &SubscribeEventServiceOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	code, err := s.SubscribeEventPort.Execute(input.UserId, input.EventId)

	if err != nil {
		if errors.Is(err, domainErrors.ErrEventNotFound) {
			return &SubscribeEventServiceOutput{
				Message:    "Esse evento não existe",
				StatusCode: 404,
			}
		}

		if errors.Is(err, domainErrors.ErrUserAlreadySubscribed) {
			return &SubscribeEventServiceOutput{
				Message:    "Você já está inscrito nesse evento",
				StatusCode: 400,
			}
		}

		return &SubscribeEventServiceOutput{
			Message:    "Não foi possível se inscrever nesse evento.",
			StatusCode: 500,
		}
	}

	return &SubscribeEventServiceOutput{
		Message:    "Inscrição confirmada",
		Code:       code,
		StatusCode: 201,
	}
}
