package usecaseImplementation

import (
	createevent "github.com/velo-project/events-service/core/use_cases/create_event"
	"github.com/velo-project/events-service/core/use_cases/create_event/implementation/ports"
	"github.com/velo-project/events-service/core/use_cases/create_event/io"
)

type CreateEventUseCase struct {
	port ports.PersistNewEventPort
}

func NewCreateEventUseCase(port ports.PersistNewEventPort) createevent.CreateEventUseCase {
	return &CreateEventUseCase{
		port: port,
	}
}

func (ceu *CreateEventUseCase) Execute(input io.CreateEventInput) io.CreateEventOutput {
	portInput := mapCreateEventInputToPersistNewEventPortInput(input)

	savedEntity := ceu.port.Execute(portInput)

	return io.CreateEventOutput{
		Message:    "success",
		Event:      savedEntity.Event,
		StatusCode: 201,
	}
}

func mapCreateEventInputToPersistNewEventPortInput(usecaseInput io.CreateEventInput) ports.PersistNewEventPortInput {
	return ports.PersistNewEventPortInput{
		Name:        usecaseInput.Name,
		Description: usecaseInput.Description,
	}
}
