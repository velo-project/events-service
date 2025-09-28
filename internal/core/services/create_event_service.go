package services

import (
	"fmt"
	"io"
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
	SaveFilePort        ports.SaveFilePort
}

type CreateEventServiceInput struct {
	Name           string    `form:"name"`
	Description    *string   `form:"description"`
	Location       *string   `form:"location"`
	Date           time.Time `form:"date"`
	Image          io.Reader `form:"-"`
	ImageExtension string    `form:"-"`
}

type CreateEventServiceOutput struct {
	Message    string `json:"message"`
	EventId    *int   `json:"event_id"`
	StatusCode int    `json:"status_code"`
}

func NewCreateEventService(ce ports.CreateEventPort, eg ports.EmbeddingsGenerator, sf ports.SaveFilePort) CreateEventService {
	return &createEventService{
		CreateEventPort:     ce,
		EmbeddingsGenerator: eg,
		SaveFilePort:        sf,
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

	if input.Image != nil {
		if !isValidExtension(input.ImageExtension) {
			return &CreateEventServiceOutput{
				Message:    "Extensão de imagem inválida",
				StatusCode: 400,
			}
		}

		imageUrl, err := s.SaveFilePort.Execute(input.Image)
		if err != nil {
			return &CreateEventServiceOutput{
				Message:    "Não foi possível salvar a imagem",
				StatusCode: 500,
			}
		}
		event.ImageURL = imageUrl
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

func isValidExtension(ext string) bool {
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
