package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"gitlab.com/velo-company/services/events-service/internal/adapters/ai"
	"gitlab.com/velo-company/services/events-service/internal/adapters/database"
	"gitlab.com/velo-company/services/events-service/internal/core/services"
)

type CreateEventHandler struct {
	db    *sql.DB
	model *genai.EmbeddingModel
}

func NewCreateEventHandler(db *sql.DB, md *genai.EmbeddingModel) *CreateEventHandler {
	return &CreateEventHandler{
		db:    db,
		model: md,
	}
}

func (h *CreateEventHandler) Handle(c *gin.Context) {
	var input services.CreateEventServiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	createEventPort := database.NewCreateEventAdapter(h.db)
	embeddingsGenerator := ai.NewEmbeddingsGenerator(h.model)
	service := services.NewCreateEventService(createEventPort, embeddingsGenerator)

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
