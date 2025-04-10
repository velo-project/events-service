package createevent

import "github.com/velo-project/events-service/core/use_cases/create_event/io"

type CreateEventUseCase interface {
	Execute(input io.CreateEventInput) io.CreateEventOutput
}
