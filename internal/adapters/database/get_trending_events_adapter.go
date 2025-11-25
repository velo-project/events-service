package database

import (
	"database/sql"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
)

type GetTrendingEventsAdapter struct {
	DB *sql.DB
}

func NewGetTrendingEventsAdapter(db *sql.DB) *GetTrendingEventsAdapter {
	return &GetTrendingEventsAdapter{DB: db}
}

func (a *GetTrendingEventsAdapter) GetTrendingEvents() ([]entities.Event, error) {
	rows, err := a.DB.Query(`
		SELECT e.id_event, e.name_event, e.description_event, e.location_event, e.photo_event, e.date_event, e.active_event, e.canceled_event, e.suspended_event
		FROM tb_events e
		JOIN tb_user_events ue ON e.id_event = ue.fk_id_event
		WHERE e.active_event = TRUE AND e.canceled_event = FALSE AND e.suspended_event = FALSE AND e.date_event > NOW() AND ue.created_at >= NOW() - INTERVAL '7 days'
		GROUP BY e.id_event
		ORDER BY COUNT(ue.id_user_event) DESC
		LIMIT 10
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []entities.Event
	for rows.Next() {
		var event entities.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Photo, &event.Date, &event.IsActive, &event.IsCanceled, &event.IsSuspended); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
