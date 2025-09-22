package services

import "gitlab.com/velo-company/services/events-service/internal/core/ports"

type CancelSubscriptionService interface {
	Execute(input *CancelSubscriptionServiceInput) *CancelSubscriptionServiceOutput
}
type CancelSubscriptionServiceInput struct {
	UserId  int
	EventId int
}
type CancelSubscriptionServiceOutput struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
type cancelSubscriptionService struct {
	csp ports.CancelSubscriptionPort
	uei ports.UserExistsByIdPort
}

func NewCancelSubscriptionService(csp ports.CancelSubscriptionPort, uei ports.UserExistsByIdPort) CancelSubscriptionService {
	return &cancelSubscriptionService{
		csp: csp,
		uei: uei,
	}
}

func (c cancelSubscriptionService) Execute(input *CancelSubscriptionServiceInput) *CancelSubscriptionServiceOutput {
	exists, err := c.uei.Execute(input.UserId)
	if err != nil {
		return &CancelSubscriptionServiceOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamento mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return &CancelSubscriptionServiceOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	err = c.csp.Execute(input.EventId, input.UserId)

	if err != nil {
		return &CancelSubscriptionServiceOutput{
			Message:    "Não foi possível cancelar sua inscrição",
			StatusCode: 500,
		}
	}

	return &CancelSubscriptionServiceOutput{
		Message:    "OK",
		StatusCode: 200,
	}
}
