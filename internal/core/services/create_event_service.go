package services

import (
	"time"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type CreateEventService interface {
	Execute(input *CreateEventServiceInput) *CreateEventServiceOutput
}

type createEventService struct {
	CreateEventPort ports.CreateEventPort
}

type CreateEventServiceInput struct {
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Location    *string   `json:"location"`
	Date        time.Time `json:"date"`
}

type CreateEventServiceOutput struct {
	Message    string `json:"message"`
	EventId    *int   `json:"event_id"`
	StatusCode int    `json:"status_code"`
}

func NewCreateEventService(ce ports.CreateEventPort) CreateEventService {
	return &createEventService{
		CreateEventPort: ce,
	}
}

func (s createEventService) Execute(input *CreateEventServiceInput) *CreateEventServiceOutput {
	// TODO: Add date parsing and validation
	event := entities.Event{
		Name:        input.Name,
		Description: input.Description,
		Location:    input.Location,
		Date:        input.Date,
	}

	eventId, err := s.CreateEventPort.Execute(&event)

	if err != nil {
		return &CreateEventServiceOutput{
			Message:    "Não foi possível criar esse evento",
			StatusCode: 500,
		}
	}

	return &CreateEventServiceOutput{
		Message:    "Evento criado com sucesso",
		EventId:    eventId,
		StatusCode: 201,
	}
}
