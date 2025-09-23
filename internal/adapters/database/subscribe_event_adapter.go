package database

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"time"

	domainErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
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
	subscribeEventQuery                  = `INSERT INTO tb_user_events (fk_id_user, fk_id_event, participation_status_event, confirmation_code_event) VALUES ($1, $2, $3, $4)`
)

func (s subscribeEventAdapter) Execute(userId int, eventId int) (*string, error) {
	tx, err := s.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var eventExists int
	err = tx.QueryRow(verifyIfEventExistsQuery, eventId).Scan(&eventExists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.ErrEventNotFound
		}
		return nil, err
	}

	var userRegistered int
	err = tx.QueryRow(verifyIfUserIsAlreadyRegisteredQuery, userId, eventId).Scan(&userRegistered)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

	}

	if userRegistered == 1 {
		return nil, domainErrors.ErrUserAlreadySubscribed
	}

	code := generateCode()
	_, err = tx.Exec(subscribeEventQuery, userId, eventId, code)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &code, nil
}

func generateCode() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits := []rune("0123456789")

	rand.New(rand.NewSource(time.Now().UnixNano()))

	var chars []rune
	for i := 0; i < 3; i++ {
		chars = append(chars, letters[rand.Intn(len(letters))])
	}

	for i := 0; i < 3; i++ {
		chars = append(chars, digits[rand.Intn(len(digits))])
	}

	rand.Shuffle(len(chars), func(i, j int) {
		chars[i], chars[j] = chars[j], chars[i]
	})

	return string(chars)
}
