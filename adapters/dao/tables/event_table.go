package tables

import (
	"database/sql"
)

type EventTable struct {
	ID          uint           `gorm:"column:event_id;primaryKey;unique;autoIncrement"`
	Name        string         `gorm:"column:event_name;not null"`
	Description sql.NullString `gorm:"column:event_description;null"`
	PhotoPath   sql.NullString `gorm:"column:event_photo;null"`
}

func (EventTable) TableName() string {
	return "tb_events"
}
