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

const (
	cancelEventQuery     = `UPDATE tb_events SET canceled_event = TRUE WHERE id_event = $1`
	searchEventByIdQuery = `SELECT id_event, name_event, description_event, location_event, photo_event, date_event, active_event, canceled_event, deleted_event FROM tb_events WHERE id_event = $1`
)

func (a *cancelEventAdapter) Execute(eventId int) error {
	row := a.Db.QueryRow(searchEventByIdQuery, eventId)

	var event entities.Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Photo, &event.Date, &event.Active, &event.Canceled, &event.Deleted)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.ErrEventNotFound
		}
		return err
	}

	_, err = a.Db.Exec(cancelEventQuery, eventId)

	return err
}
