package services

import (
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type SubscribeEventService interface {
	Execute(input *SubscribeEventServiceInput) (*SubscribeEventServiceOutput, error)
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
	Message string  `json:"message"`
	Code    *string `json:"code"`
}

func NewSubscribeEventService(ue ports.UserExistsByIdPort, se ports.SubscribeEventPort) SubscribeEventService {
	return &subscribeEventService{
		UserExistsByIdPort: ue,
		SubscribeEventPort: se,
	}
}

func (s subscribeEventService) Execute(input *SubscribeEventServiceInput) (*SubscribeEventServiceOutput, error) {
	exists, err := s.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Este usuário não existe")
	}

	code, err := s.SubscribeEventPort.Execute(input.UserId, input.EventId)

	if err != nil {
		return nil, err
	}

	return &SubscribeEventServiceOutput{
		Message: "Inscrição confirmada",
		Code:    code,
	}, nil
}
