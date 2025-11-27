package services

import (
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type ActivateEventService interface {
	Execute(input *ActivateEventServiceInput) (*ActivateEventServiceOutput, error)
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
	Message string `json:"message"`
}

func (s activateEventService) Execute(input *ActivateEventServiceInput) (*ActivateEventServiceOutput, error) {
	exists, err := s.userExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Este usuário não existe")
	}

	err = s.activateEventPort.Execute(input.EventId)
	if err != nil {
		return nil, err
	}

	return &ActivateEventServiceOutput{
		Message: "OK",
	}, nil
}
