package database

import (
	"database/sql"
	errors "errors"
	"time"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	domainErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type cancelSubscriptionAdapter struct {
	DB *sql.DB
}

func NewCancelSubscriptionAdapter(db *sql.DB) ports.CancelSubscriptionPort {
	return &cancelSubscriptionAdapter{
		DB: db,
	}
}

const (
	getEventDateQuery = `SELECT date_event FROM tb_events WHERE id_event = $1`
	searchEventQuery  = `SELECT 1 FROM tb_user_events WHERE fk_id_event = $1 AND fk_id_user = $2`
	cancelEventQuery  = `UPDATE tb_user_events SET participation_status_event = $1 WHERE fk_id_event = $2 AND fk_id_user = $3`
)

func (c cancelSubscriptionAdapter) Execute(eventId int, userId int) error {
	var eventExists int
	err := c.DB.QueryRow(searchEventQuery, eventId, userId).Scan(&eventExists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.ErrUserSubscriptionNotFound
		}
		return err
	}

	var eventDate time.Time
	err = c.DB.QueryRow(getEventDateQuery, eventId).Scan(&eventDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domainErrors.ErrEventNotFound
		}
		return err
	}

	if time.Until(eventDate).Hours()/24 <= 7 {
		return domainErrors.ErrBlockedCancelSubscription
	}

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(cancelEventQuery, entities.Cancelled, eventId, userId)

	if err != nil {
		return err
	}

	return tx.Commit()
}
