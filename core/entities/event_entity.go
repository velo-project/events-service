package entities

import (
	"time"

	"github.com/velo-project/events-service/core/entities/enums"
)

type EventEntity struct {
	ID                   int                  `json:"id"`
	Name                 string               `json:"name"`
	Description          string               `json:"description"`
	PhotoPath            string               `json:"photoPath"`
	IsActive             bool                 `json:"isActive"`
	IsCanceled           bool                 `json:"isCanceled"`
	Address              string               `json:"address"`
	DifficultyOfTheTrack enums.DifficultyEnum `json:"difficultyOfTheTrack"`
	TrackType            enums.TrackTypeEnum  `json:"trackType"`
	WhenWillHappen       time.Time            `json:"whenWillHappen"`
	CreatedAt            time.Time            `json:"createdAt"`
	UpdatedAt            time.Time            `json:"updatedAt"`
}

func (e *EventEntity) Cancel() {
	if e.IsCanceled {
		return
	}
	e.IsCanceled = true
	e.UpdatedAt = time.Now()
}

func (e *EventEntity) Activate() {
	if e.IsActive {
		return
	}
	e.IsActive = true
	e.UpdatedAt = time.Now()
}

func (e *EventEntity) Desactivate() {
	if !e.IsActive {
		return
	}
	e.IsActive = false
	e.UpdatedAt = time.Now()
}

func (e *EventEntity) IsUpcoming() bool {
	return time.Now().Before(e.WhenWillHappen)
}
