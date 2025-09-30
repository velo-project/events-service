package database

import (
	"database/sql"
	"errors"
	"time"

	domainErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type activateEventAdapter struct {
	Db *sql.DB
}

func NewActivateEventAdapter(db *sql.DB) ports.ActivateEventPort {
	return &activateEventAdapter{Db: db}
}

const (
	activateEventQuery = `UPDATE tb_events SET suspended_event = FALSE WHERE id_event = $1`
)

func (a *activateEventAdapter) Execute(eventId int) error {
	var eventDate time.Time
	err := a.Db.QueryRow(getEventDateQuery, eventId).Scan(&eventDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.ErrEventNotFound
		}
		return err
	}

	if time.Until(eventDate).Hours()/24 <= 3 {
		return domainErrors.ErrBlockedActivateEvent
	}

	_, err = a.Db.Exec(activateEventQuery, eventId)

	return err
}
