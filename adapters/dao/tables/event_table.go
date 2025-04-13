package tables

import (
	"database/sql"
	"time"
)

type EventTable struct {
	ID                   uint           `gorm:"column:event_id;primaryKey;unique;autoIncrement"`
	Name                 string         `gorm:"column:event_name;not null"`
	Description          sql.NullString `gorm:"column:event_description;null"`
	PhotoPath            sql.NullString `gorm:"column:event_photo;null"`
	IsActive             int            `gorm:"column:event_active;not null"`
	IsCanceled           int            `gorm:"column:event_canceled;not null"`
	Address              string         `gorm:"column:event_address;not null"`
	DifficultyOfTheTrack int            `gorm:"column:event_difficulty;not null"`
	TrackType            int            `gorm:"column:event_track_type;not null"`
	WhenWillHappen       time.Time      `gorm:"column:event_when_will_happen;not null"`
	CreatedAt            time.Time      `gorm:"column:event_created_at;not null"`
	UpdatedAt            time.Time      `gorm:"column:event_updated_at;not null"`
}

func (EventTable) TableName() string {
	return "tb_events"
}
