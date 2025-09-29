package database

import (
	"database/sql"
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	domainErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type cancelEventAdapter struct {
	Db *sql.DB
}

func NewCancelEventAdapter(db *sql.DB) ports.CancelEventPort {
	return &cancelEventAdapter{Db: db}
}

func (a *cancelEventAdapter) CancelEvent(eventId int) error {
	query := "UPDATE tb_events SET canceled_event = TRUE WHERE id_event = $1"

	_, err := a.Db.Exec(query, eventId)

	return err
}

func (a *cancelEventAdapter) GetEventById(eventId int) (*entities.Event, error) {
	query := "SELECT id_event, name_event, description_event, location_event, photo_event, date_event, active_event, canceled_event, deleted_event FROM tb_events WHERE id_event = $1"

	row := a.Db.QueryRow(query, eventId)

	var event entities.Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Photo, &event.Date, &event.Active, &event.Canceled, &event.Deleted)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.ErrEventNotFound
		}
		return nil, err
	}

	return &event, nil
}
