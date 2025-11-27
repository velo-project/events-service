package services

import (
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type CancelEventService interface {
	Execute(input *CancelEventServiceInput) (*CancelEventServiceOutput, error)
}

type cancelEventService struct {
	cancelEventPort    ports.CancelEventPort
	userExistsByIdPort ports.UserExistsByIdPort
}

func NewCancelEventService(cancelEventPort ports.CancelEventPort, userExistsByIdPort ports.UserExistsByIdPort) CancelEventService {
	return &cancelEventService{cancelEventPort: cancelEventPort, userExistsByIdPort: userExistsByIdPort}
}

type CancelEventServiceInput struct {
	EventId int
	UserId  int
}

type CancelEventServiceOutput struct {
	Message string `json:"message"`
}

func (s cancelEventService) Execute(input *CancelEventServiceInput) (*CancelEventServiceOutput, error) {
	exists, err := s.userExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Este usuário não existe")
	}

	err = s.cancelEventPort.Execute(input.EventId)
	if err != nil {
		return nil, err
	}

	return &CancelEventServiceOutput{
		Message: "OK",
	}, nil
}
