package database

import (
	"database/sql"

	"gitlab.com/velo-company/services/events-service/internal/core/entities"
)

type GetSubscribedEventsAdapter struct {
	DB *sql.DB
}

func NewGetSubscribedEventsAdapter(db *sql.DB) *GetSubscribedEventsAdapter {
	return &GetSubscribedEventsAdapter{DB: db}
}

func (a *GetSubscribedEventsAdapter) GetSubscribedEvents(userID string) ([]entities.Event, error) {
	rows, err := a.DB.Query(`
		SELECT e.id_event, e.name_event, e.description_event, e.location_event, e.photo_event, e.date_event, e.active_event, e.canceled_event, e.suspended_event
		FROM tb_events e
		JOIN tb_user_events ue ON e.id_event = ue.fk_id_event
		WHERE ue.fk_id_user = $1 AND ue.participation_status_event = 0 AND e.date_event > NOW()
		ORDER BY e.date_event ASC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []entities.Event
	for rows.Next() {
		var event entities.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Photo, &event.Date, &event.Active, &event.Canceled, &event.Suspended); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
