package services

import (
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type SuspendEventService interface {
	Execute(input *SuspendEventServiceInput) (*SuspendEventServiceOutput, error)
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
	Message string `json:"message"`
}

func (s suspendEventService) Execute(input *SuspendEventServiceInput) (*SuspendEventServiceOutput, error) {
	exists, err := s.userExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Este usuário não existe")
	}

	err = s.suspendEventPort.Execute(input.EventId)
	if err != nil {
		return nil, err
	}

	return &SuspendEventServiceOutput{
		Message: "OK",
	}, nil
}
