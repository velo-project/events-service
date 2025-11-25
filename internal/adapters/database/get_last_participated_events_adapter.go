package database

import (
	"database/sql"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
)

type GetLastParticipatedEventsAdapter struct {
	DB *sql.DB
}

func NewGetLastParticipatedEventsAdapter(db *sql.DB) *GetLastParticipatedEventsAdapter {
	return &GetLastParticipatedEventsAdapter{DB: db}
}

func (a *GetLastParticipatedEventsAdapter) GetLastParticipatedEvents(userID string) ([]entities.Event, error) {
	rows, err := a.DB.Query(`
		SELECT e.id_event, e.name_event, e.description_event, e.location_event, e.photo_event, e.date_event, e.active_event, e.canceled_event, e.suspended_event
		FROM tb_events e
		JOIN tb_user_events ue ON e.id_event = ue.fk_id_event
		WHERE ue.fk_id_user = $1 AND ue.participation_status_event = 1
		ORDER BY e.date_event DESC
		LIMIT 10
	`, userID)
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
