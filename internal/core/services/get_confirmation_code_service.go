package services

import (
	"time"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type GetConfirmationCodeService interface {
	Execute(input *GetConfirmationCodeInput) *GetConfirmationCodeOutput
}

type GetConfirmationCodeInput struct {
	UserId  int
	EventId int
}

type GetConfirmationCodeOutput struct {
	Code       *string `json:"code"`
	Message    string  `json:"message"`
	StatusCode int     `json:"status_code"`
}

type getConfirmationCodeService struct {
	GetConfirmationCodePort ports.GetConfirmationCodePort
	UserExistsByIdPort      ports.UserExistsByIdPort
}

func NewGetConfirmationCodeService(gc ports.GetConfirmationCodePort, up ports.UserExistsByIdPort) GetConfirmationCodeService {
	return &getConfirmationCodeService{
		GetConfirmationCodePort: gc,
		UserExistsByIdPort:      up,
	}
}

func (s *getConfirmationCodeService) Execute(input *GetConfirmationCodeInput) *GetConfirmationCodeOutput {
	exists, err := s.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return &GetConfirmationCodeOutput{
			Message:    "Estamos enfrentando problemas no momento. Tente novamente mais tarde",
			StatusCode: 502,
		}
	}
	if !exists {
		return &GetConfirmationCodeOutput{
			Message:    "Este usuário não existe",
			StatusCode: 404,
		}
	}

	code, eventDate, err := s.GetConfirmationCodePort.Execute(input.UserId, input.EventId)
	if err != nil {
		return &GetConfirmationCodeOutput{
			Message:    "Não foi possível buscar o código de confirmação",
			StatusCode: 500,
		}
	}

	if code == nil {
		return &GetConfirmationCodeOutput{
			Message:    "Inscrição não encontrada para este evento",
			StatusCode: 404,
		}
	}

	if eventDate.Before(time.Now()) {
		return &GetConfirmationCodeOutput{
			Message:    "Este evento já ocorreu",
			StatusCode: 410, // Gone
		}
	}

	return &GetConfirmationCodeOutput{
		Code:       code,
		Message:    "OK",
		StatusCode: 200,
	}
}
