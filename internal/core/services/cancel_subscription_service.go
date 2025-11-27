package services

import (
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type CancelSubscriptionService interface {
	Execute(input *CancelSubscriptionServiceInput) (*CancelSubscriptionServiceOutput, error)
}
type CancelSubscriptionServiceInput struct {
	UserId  int
	EventId int
}
type CancelSubscriptionServiceOutput struct {
	Message string `json:"message"`
}
type cancelSubscriptionService struct {
	CancelSubscriptionPort ports.CancelSubscriptionPort
	UserExistsByIdPort     ports.UserExistsByIdPort
}

func NewCancelSubscriptionService(CancelSubscriptionPort ports.CancelSubscriptionPort, UserExistsByIdPort ports.UserExistsByIdPort) CancelSubscriptionService {
	return &cancelSubscriptionService{
		CancelSubscriptionPort: CancelSubscriptionPort,
		UserExistsByIdPort:     UserExistsByIdPort,
	}
}

func (c cancelSubscriptionService) Execute(input *CancelSubscriptionServiceInput) (*CancelSubscriptionServiceOutput, error) {
	exists, err := c.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Este usuário não existe")
	}

	err = c.CancelSubscriptionPort.Execute(input.EventId, input.UserId)

	if err != nil {
		return nil, err
	}

	return &CancelSubscriptionServiceOutput{
		Message: "OK",
	}, nil
}
