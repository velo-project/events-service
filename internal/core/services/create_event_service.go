package services

import (
	"errors"
	"fmt"
	"io"
	"time"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	customErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type CreateEventService interface {
	Execute(input *CreateEventServiceInput) (*CreateEventServiceOutput, error)
}

type createEventService struct {
	CreateEventPort     ports.CreateEventPort
	EmbeddingsGenerator ports.EmbeddingsGenerator
	SaveFilePort        ports.SaveFilePort
	UserExistsByIdPort  ports.UserExistsByIdPort
}

type CreateEventServiceInput struct {
	Name           string    `form:"name"`
	Description    *string   `form:"description"`
	Location       *string   `form:"location"`
	Date           time.Time `form:"date"`
	Image          io.Reader `form:"-"`
	ImageExtension string    `form:"-"`
	UserId         int       `form:"-"`
}

type CreateEventServiceOutput struct {
	Message    string `json:"message"`
	EventId    *int   `json:"event_id"`
	StatusCode int    `json:"-"`
}

func NewCreateEventService(ce ports.CreateEventPort, eg ports.EmbeddingsGenerator, sf ports.SaveFilePort, ue ports.UserExistsByIdPort) CreateEventService {
	return &createEventService{
		CreateEventPort:     ce,
		EmbeddingsGenerator: eg,
		SaveFilePort:        sf,
		UserExistsByIdPort:  ue,
	}
}

func (s createEventService) Execute(input *CreateEventServiceInput) (*CreateEventServiceOutput, error) {
	exists, err := s.UserExistsByIdPort.Execute(input.UserId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Este usuário não existe")
	}
	// TODO: Add date parsing and validation
	textToEmbeddings := fmt.Sprintf("%s %s", input.Name, deref(input.Description))
	embeddings, err := s.EmbeddingsGenerator.Generate(ports.EmbeddingsGeneratorInput{
		Text: textToEmbeddings,
	})

	if err != nil {
		return nil, customErrors.ErrEventNotCreated
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
			return nil, errors.New("Extensão de imagem inválida")
		}

		imageUrl, err := s.SaveFilePort.Execute(input.Image)
		if err != nil {
			return nil, customErrors.ErrEventNotCreated
		}
		event.ImageURL = imageUrl
	}

	eventId, err := s.CreateEventPort.Execute(&event)

	if err != nil {
		return nil, customErrors.ErrEventNotCreated
	}

	return &CreateEventServiceOutput{
		Message: "Evento criado com sucesso",
		EventId: eventId,
	}, nil
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
