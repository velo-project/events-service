package database

import (
	"database/sql"
	"errors"
	"time"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type getConfirmationCodeAdapter struct {
	DB *sql.DB
}

func NewGetConfirmationCodeAdapter(DB *sql.DB) ports.GetConfirmationCodePort {
	return &getConfirmationCodeAdapter{
		DB: DB,
	}
}

const getConfirmationCodeQuery = `
	SELECT ue.confirmation_code_event, e.date
	FROM tb_user_events ue
	JOIN tb_events e ON ue.fk_id_event = e.id
	WHERE ue.fk_id_user = $1 AND ue.fk_id_event = $2`

func (g *getConfirmationCodeAdapter) Execute(userId int, eventId int) (*string, *time.Time, error) {
	var code sql.NullString
	var eventDate time.Time

	err := g.DB.QueryRow(getConfirmationCodeQuery, userId, eventId).Scan(&code, &eventDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, nil // No subscription found
		}
		return nil, nil, err
	}

	if !code.Valid {
		return nil, nil, nil // Code is null
	}

	return &code.String, &eventDate, nil
}
