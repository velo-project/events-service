package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/events-service/internal/adapters/database"
	"gitlab.com/velo-company/services/events-service/internal/core/services"
)

type CreateEventHandler struct {
	db *sql.DB
}

func NewCreateEventHandler(db *sql.DB) *CreateEventHandler {
	return &CreateEventHandler{db: db}
}

func (h *CreateEventHandler) Handle(c *gin.Context) {
	var input services.CreateEventServiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	createEventPort := database.NewCreateEventAdapter(h.db)
	service := services.NewCreateEventService(createEventPort)

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
