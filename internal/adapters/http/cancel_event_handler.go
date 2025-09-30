package http

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/events-service/internal/adapters/database"
	"gitlab.com/velo-company/services/events-service/internal/core/services"
)

type CancelEventHandler struct {
	db *sql.DB
}

func NewCancelEventHandler(db *sql.DB) *CancelEventHandler {
	return &CancelEventHandler{db: db}
}

// @Summary Cancel an event
// @Description Cancel an event
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} services.CancelEventServiceOutput
// @Failure 400 {object} services.CancelEventServiceOutput
// @Failure 404 {object} services.CancelEventServiceOutput
// @Failure 500 {object} services.CancelEventServiceOutput
// @Security Bearer
// @Router /cancel/{id} [patch]
func (h *CancelEventHandler) Handle(c *gin.Context) {
	eventIdStr := c.Param("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "ID de evento inválido", "status_code": 400})
		return
	}

	cancelEventAdapter := database.NewCancelEventAdapter(h.db)
	service := services.NewCancelEventService(cancelEventAdapter)

	input := services.CancelEventServiceInput{
		EventId: eventId,
	}

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
