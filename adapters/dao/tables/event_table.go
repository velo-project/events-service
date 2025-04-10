package tables

import (
	"database/sql"

	"gorm.io/gorm"
)

type EventTable struct {
	gorm.Model
	ID          uint           `gorm:"primaryKey;unique;autoIncrement"`
	Name        string         `gorm:"not null"`
	Description sql.NullString `gorm:"null"`
	PhotoPath   sql.NullString `gorm:"null"`
}

func (EventTable) TableName() string {
	return "events"
}
