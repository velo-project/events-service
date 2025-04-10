package ports

import "github.com/velo-project/events-service/core/entities"

type PersistNewEventPort interface {
	Execute(input PersistNewEventPortInput) PersistNewEventPortOutput
}

type PersistNewEventPortInput struct {
	Name        string
	Description string
}

type PersistNewEventPortOutput struct {
	Event entities.EventEntity
}
