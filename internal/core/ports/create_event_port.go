package ports

import "gitlab.com/velo-company/services/events-service/internal/core/entities"

type CreateEventPort interface {
	Execute(event *entities.Event) (*int, error)
}
