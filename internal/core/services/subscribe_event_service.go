package services

import "gitlab.com/velo-company/services/events-service/internal/core/ports"

type SubscribeEventService interface {
	Execute(input SubscribeEventServiceInput) SubscribeEventServiceOutput
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
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewSubscribeEventService(ue ports.UserExistsByIdPort, se ports.SubscribeEventPort) SubscribeEventService {
	return &subscribeEventService{
		UserExistsByIdPort: ue,
		SubscribeEventPort: se,
	}
}

func (s subscribeEventService) Execute(input SubscribeEventServiceInput) SubscribeEventServiceOutput {
	exists, err := s.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return SubscribeEventServiceOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return SubscribeEventServiceOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	err = s.SubscribeEventPort.Execute(input.UserId, input.EventId)

	if err != nil {
		return SubscribeEventServiceOutput{
			Message:    "Não foi possível se inscrever nesse evento.",
			StatusCode: 500,
		}
	}

	return SubscribeEventServiceOutput{
		Message:    "Inscrição confirmada",
		StatusCode: 201,
	}
}
