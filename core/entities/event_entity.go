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
