package io

import (
	"github.com/velo-project/events-service/core/entities"
)

type CreateEventOutput struct {
	Message    string               `json:"message"`
	Event      entities.EventEntity `json:"event"`
	StatusCode int                  `json:"statusCode"`
}
