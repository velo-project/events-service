package ports

import "gitlab.com/velo-company/services/events-service/internal/core/entities"

type CancelEventPort interface {
	CancelEvent(eventId int) error
	GetEventById(eventId int) (*entities.Event, error)
}
