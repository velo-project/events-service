package database

import (
	"database/sql"

	"github.com/pgvector/pgvector-go"
	"gitlab.com/velo-company/services/events-service/internal/core/entities"
)

type GetRecommendedEventsAdapter struct {
	DB *sql.DB
}

func NewGetRecommendedEventsAdapter(db *sql.DB) *GetRecommendedEventsAdapter {
	return &GetRecommendedEventsAdapter{DB: db}
}

func (a *GetRecommendedEventsAdapter) GetRecommendedEvents(userID string) ([]entities.Event, error) {
	// 1. Get user's participated events' embeddings
	rows, err := a.DB.Query(`
		SELECT e.embeddings_event
		FROM tb_events e
		JOIN tb_user_events ue ON e.id_event = ue.fk_id_event
		WHERE ue.fk_id_user = $1 AND ue.participation_status_event = 1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var embeddings []pgvector.Vector
	for rows.Next() {
		var embedding pgvector.Vector
		if err := rows.Scan(&embedding); err != nil {
			return nil, err
		}
		embeddings = append(embeddings, embedding)
	}

	if len(embeddings) == 0 {
		return []entities.Event{}, nil
	}

	// 2. Calculate user preference embedding (average)
	avgEmbedding := make([]float32, 1536)
	for _, embedding := range embeddings {
		for i, val := range embedding.Slice() {
			avgEmbedding[i] += val
		}
	}
	for i := range avgEmbedding {
		avgEmbedding[i] /= float32(len(embeddings))
	}

	// 3. Find similar events
	similarEventsRows, err := a.DB.Query(`
		SELECT id_event, name_event, description_event, location_event, photo_event, date_event, active_event, canceled_event, suspended_event
		FROM tb_events
		WHERE active_event = TRUE AND canceled_event = FALSE AND suspended_event = FALSE AND date_event > NOW()
		ORDER BY embeddings_event <-> $1
		LIMIT 10
	`, pgvector.NewVector(avgEmbedding))
	if err != nil {
		return nil, err
	}
	defer similarEventsRows.Close()

	var events []entities.Event
	for similarEventsRows.Next() {
		var event entities.Event
		if err := similarEventsRows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Photo, &event.Date, &event.Active, &event.Canceled, &event.Suspended); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
