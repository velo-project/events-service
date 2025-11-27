package services

import (
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type ConfirmSubscriptionService interface {
	Execute(input *ConfirmSubscriptionInput) (*ConfirmSubscriptionOutput, error)
}

type ConfirmSubscriptionInput struct {
	Code    string
	UserId  int
	EventId int
}

type ConfirmSubscriptionOutput struct {
	Message string `json:"message"`
}

type confirmSubscriptionService struct {
	ConfirmSubscriptionPort ports.ConfirmSubscriptionPort
	UserExistsByIdPort      ports.UserExistsByIdPort
}

func NewConfirmSubscriptionService(cp ports.ConfirmSubscriptionPort, up ports.UserExistsByIdPort) ConfirmSubscriptionService {
	return &confirmSubscriptionService{
		ConfirmSubscriptionPort: cp,
		UserExistsByIdPort:      up,
	}
}

func (s *confirmSubscriptionService) Execute(input *ConfirmSubscriptionInput) (*ConfirmSubscriptionOutput, error) {
	exists, err := s.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Este usuário não existe")
	}

	err = s.ConfirmSubscriptionPort.Execute(input.Code, input.UserId, input.EventId)

	if err != nil {
		return nil, err
	}

	return &ConfirmSubscriptionOutput{
		Message: "OK",
	}, nil
}
