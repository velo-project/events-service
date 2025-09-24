package database

import (
	"database/sql"
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type confirmSubscriptionAdapter struct {
	DB *sql.DB
}

func NewConfirmSubscriptionAdapter(DB *sql.DB) ports.ConfirmSubscriptionPort {
	return &confirmSubscriptionAdapter{
		DB: DB,
	}
}

const (
	verifyUserCodeQuery       = `SELECT tb_user_events FROM tb_user_events WHERE user_id = $1 AND event_id = $2`
	confirmParticipationQuery = `UPDATE tb_user_events SET participation_status_event = $1 WHERE fk_id_user = $2 AND fk_id_event = $3`
)

func (c confirmSubscriptionAdapter) Execute(code string, userId int, eventId int) error {
	var confirmationCode string

	err := c.DB.QueryRow(verifyUserCodeQuery, userId, eventId).Scan(&confirmationCode)
	if err != nil {
		return err
	}

	if confirmationCode != code {
		return domainErrors.ErrInvalidConfirmationCode
	}

	_, err = c.DB.Exec(confirmParticipationQuery, entities.Participated, userId, eventId)
	if err != nil {
		return err
	}

	return nil
}
