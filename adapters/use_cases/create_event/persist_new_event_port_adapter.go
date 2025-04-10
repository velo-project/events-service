package createevent

import (
	"github.com/velo-project/events-service/adapters/dao/repositories"
	"github.com/velo-project/events-service/core/entities"
	"github.com/velo-project/events-service/core/use_cases/create_event/implementation/ports"
)

type PersistNewEventPortAdapter struct {
	repository repositories.EventRepository
}

func NewPersistNewEventPortAdapter(repository repositories.EventRepository) ports.PersistNewEventPort {
	return &PersistNewEventPortAdapter{
		repository: repository,
	}
}

func (pnepa *PersistNewEventPortAdapter) Execute(input ports.PersistNewEventPortInput) ports.PersistNewEventPortOutput {
	eventEntity := mapPersistNewEventPortInputToEventEntity(input)

	result := pnepa.repository.CreateEvent(eventEntity)

	return mapEventEntityToPersistNewEventPortOutput(result)
}

func mapPersistNewEventPortInputToEventEntity(portInput ports.PersistNewEventPortInput) entities.EventEntity {
	return entities.EventEntity{
		Name:        portInput.Name,
		Description: portInput.Description,
	}
}

func mapEventEntityToPersistNewEventPortOutput(eventEntity entities.EventEntity) ports.PersistNewEventPortOutput {
	return ports.PersistNewEventPortOutput{
		Event: eventEntity,
	}
}
