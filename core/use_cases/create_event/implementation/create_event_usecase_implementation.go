package usecaseImplementation

import (
	"github.com/velo-project/events-service/adapters/dao/repositories"
	"github.com/velo-project/events-service/core/entities"
	createevent "github.com/velo-project/events-service/core/use_cases/create_event"
	"github.com/velo-project/events-service/core/use_cases/create_event/io"
)

type CreateEventUseCase struct {
	repository repositories.EventRepository
}

func NewCreateEventUseCase(repository repositories.EventRepository) createevent.CreateEventUseCase {
	return &CreateEventUseCase{
		repository: repository,
	}
}

func (ceu *CreateEventUseCase) Execute(input io.CreateEventInput) io.CreateEventOutput {
	event := entities.EventEntity{
		Name:        input.Name,
		Description: input.Description,
		PhotoPath:   "",
	}

	savedEntity := ceu.repository.CreateEvent(event)

	return io.CreateEventOutput{
		Message:    "success",
		Event:      savedEntity,
		StatusCode: 201,
	}
}
