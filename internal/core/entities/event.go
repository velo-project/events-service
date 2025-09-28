package entities

import "time"

type Event struct {
	ID          *int      `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Location    *string   `json:"location"`
	Photo       *string   `json:"photo"`
	Date        time.Time `json:"date"`
	Embeddings  []float32 `json:"embeddings"`
	ImageURL    *string   `json:"image_url"`
	Active      bool      `json:"active"`
	Canceled    bool      `json:"canceled"`
	Deleted     bool      `json:"deleted"`
}
