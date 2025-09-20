package database

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type subscribeEventAdapter struct {
	DB *sql.DB
}

func NewSubscribeEventAdapter(DB *sql.DB) ports.SubscribeEventPort {
	return &subscribeEventAdapter{
		DB: DB,
	}
}

const (
	verifyIfEventExistsQuery             = `SELECT 1 FROM tb_events WHERE id_event = $1`
	verifyIfUserIsAlreadyRegisteredQuery = `SELECT 1 FROM tb_user_events WHERE fk_id_user = $1 AND fk_id_event = $2`
	subscribeEventQuery                  = `INSERT INTO tb_user_events (fk_id_user, fk_id_event, participation_status_event) VALUES ($1, $2, $2)`
)

func (s subscribeEventAdapter) Execute(userId int, eventId int) error {
	tx, err := s.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var eventExists int
	err = tx.QueryRow(verifyIfEventExistsQuery, eventId).Scan(&eventExists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("event does not exist")
		}
		return err
	}

	var userRegistered int
	err = tx.QueryRow(verifyIfUserIsAlreadyRegisteredQuery, userId, eventId).Scan(&userRegistered)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}

	}

	if userRegistered == 1 {
		return errors.New("user is already subscribed to this event")
	}

	_, err = tx.Exec(subscribeEventQuery, userId, eventId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
