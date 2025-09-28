package database

import (
	"context"
	"database/sql"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type createEventAdapter struct {
	DB *sql.DB
}

func NewCreateEventAdapter(DB *sql.DB) ports.CreateEventPort {
	return &createEventAdapter{
		DB: DB,
	}
}

const (
	createEventQuery = `INSERT INTO tb_events (name_event, description_event, location_event, date_event, embeddings_event, image_url_event) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_event`
)

func (s createEventAdapter) Execute(event *entities.Event) (*int, error) {
	tx, err := s.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var eventId int
	err = tx.QueryRow(createEventQuery, event.Name, event.Description, event.Location, event.Date, event.Embeddings, event.ImageURL).Scan(&eventId)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &eventId, nil
}
