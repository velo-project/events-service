package http

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/events-service/internal/adapters/database"
	grpcadapter "gitlab.com/velo-company/services/events-service/internal/adapters/grpc"
	"gitlab.com/velo-company/services/events-service/internal/core/services"
	"google.golang.org/grpc"
)

type CancelEventHandler struct {
	db   *sql.DB
	grpc *grpc.ClientConn
}

func NewCancelEventHandler(db *sql.DB, grpc *grpc.ClientConn) *CancelEventHandler {
	return &CancelEventHandler{
		db:   db,
		grpc: grpc,
	}
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

	anonymousUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"message": "Usuário não autenticado", "status_code": 401})
		return
	}

	userId := anonymousUserId.(int)

	cancelEventAdapter := database.NewCancelEventAdapter(h.db)
	userExistsByIdPort := grpcadapter.NewUserExistsByIdAdapter(h.grpc)
	service := services.NewCancelEventService(cancelEventAdapter, userExistsByIdPort)

	input := services.CancelEventServiceInput{
		EventId: eventId,
		UserId:  userId,
	}

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
