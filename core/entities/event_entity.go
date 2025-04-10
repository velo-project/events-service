package entities

type EventEntity struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PhotoPath   string `json:"photoPath"`
}
