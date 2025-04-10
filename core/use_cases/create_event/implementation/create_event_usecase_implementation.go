package implementation

import (
	"github.com/velo-project/events-service/core/entities"
	"github.com/velo-project/events-service/core/use_cases/create_event/io"
)

type CreateEventUseCase struct {
}

func NewCreateEventUseCase() CreateEventUseCase {
	return CreateEventUseCase{}
}

func (ceu *CreateEventUseCase) Execute(input io.CreateEventInput) io.CreateEventOutput {
	event := entities.EventEntity{
		Name:        input.Name,
		Description: input.Description,
		PhotoPath:   "",
	}

	return io.CreateEventOutput{
		Message:    "success",
		Event:      event,
		StatusCode: 201,
	}
}
