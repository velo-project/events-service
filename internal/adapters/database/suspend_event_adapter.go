package database

import (
	"database/sql"
	"errors"
	"time"

	domainErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type suspendEventAdapter struct {
	Db *sql.DB
}

func NewSuspendEventAdapter(db *sql.DB) ports.SuspendEventPort {
	return &suspendEventAdapter{Db: db}
}

const (
	suspendEventQuery = `UPDATE tb_events SET suspended_event = TRUE WHERE id_event = $1`
)

func (a *suspendEventAdapter) Execute(eventId int) error {
	var eventDate time.Time
	err := a.Db.QueryRow(getEventDateQuery, eventId).Scan(&eventDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.ErrEventNotFound
		}
		return err
	}

	if time.Until(eventDate).Hours()/24 <= 3 {
		return domainErrors.ErrBlockedSuspendEvent
	}

	_, err = a.Db.Exec(suspendEventQuery, eventId)

	return err
}
