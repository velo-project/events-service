package services

import (
	"errors"
	"time"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type GetConfirmationCodeService interface {
	Execute(input *GetConfirmationCodeInput) (*GetConfirmationCodeOutput, error)
}

type GetConfirmationCodeInput struct {
	UserId  int
	EventId int
}

type GetConfirmationCodeOutput struct {
	Code *string `json:"code"`
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

func (s *getConfirmationCodeService) Execute(input *GetConfirmationCodeInput) (*GetConfirmationCodeOutput, error) {
	exists, err := s.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Este usuário não existe")
	}

	code, eventDate, err := s.GetConfirmationCodePort.Execute(input.UserId, input.EventId)
	if err != nil {
		return nil, errors.New("Não foi possível buscar o código de confirmação")
	}

	if code == nil {
		return nil, errors.New("Inscrição não encontrada para este evento")
	}

	if eventDate.Before(time.Now()) {
		return nil, errors.New("Este evento já ocorreu")
	}

	return &GetConfirmationCodeOutput{
		Code: code,
	}, nil
}
