package services

import (
	"fmt"
	"time"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type CreateEventService interface {
	Execute(input *CreateEventServiceInput) *CreateEventServiceOutput
}

type createEventService struct {
	CreateEventPort     ports.CreateEventPort
	EmbeddingsGenerator ports.EmbeddingsGenerator
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

func NewCreateEventService(ce ports.CreateEventPort, eg ports.EmbeddingsGenerator) CreateEventService {
	return &createEventService{
		CreateEventPort:     ce,
		EmbeddingsGenerator: eg,
	}
}

func (s createEventService) Execute(input *CreateEventServiceInput) *CreateEventServiceOutput {
	// TODO: Add date parsing and validation
	textToEmbeddings := fmt.Sprintf("%s %s", input.Name, deref(input.Description))
	embeddings, err := s.EmbeddingsGenerator.Generate(ports.EmbeddingsGeneratorInput{
		Text: textToEmbeddings,
	})

	if err != nil {
		return &CreateEventServiceOutput{
			Message:    "Não foi possível criar esse evento",
			StatusCode: 500,
		}
	}

	event := entities.Event{
		Name:        input.Name,
		Description: input.Description,
		Location:    input.Location,
		Date:        input.Date,
		Embeddings:  embeddings.Values,
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

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
