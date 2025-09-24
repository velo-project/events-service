package services

import (
	"errors"

	domainErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type ConfirmSubscriptionService interface {
	Execute(input *ConfirmSubscriptionInput) *ConfirmSubscriptionOutput
}

type ConfirmSubscriptionInput struct {
	Code    string
	UserId  int
	EventId int
}

type ConfirmSubscriptionOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
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

func (s *confirmSubscriptionService) Execute(input *ConfirmSubscriptionInput) *ConfirmSubscriptionOutput {
	exists, err := s.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return &ConfirmSubscriptionOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamente mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return &ConfirmSubscriptionOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	err = s.ConfirmSubscriptionPort.Execute(input.Code, input.UserId, input.EventId)

	if err != nil {
		if errors.Is(err, domainErrors.ErrInvalidConfirmationCode) {
			return &ConfirmSubscriptionOutput{
				Message:    err.Error(),
				StatusCode: 400,
			}
		}
		return &ConfirmSubscriptionOutput{
			Message:    "Não foi possível confirmar sua inscrição",
			StatusCode: 500,
		}
	}

	return &ConfirmSubscriptionOutput{
		Message:    "OK",
		StatusCode: 200,
	}
}
