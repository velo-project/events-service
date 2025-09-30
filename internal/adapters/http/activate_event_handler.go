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

type ActivateEventHandler struct {
	db   *sql.DB
	grpc *grpc.ClientConn
}

func NewActivateEventHandler(db *sql.DB, grpc *grpc.ClientConn) *ActivateEventHandler {
	return &ActivateEventHandler{db: db, grpc: grpc}
}

func (h *ActivateEventHandler) Handle(c *gin.Context) {
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

	activateEventAdapter := database.NewActivateEventAdapter(h.db)
	userExistsByIdPort := grpcadapter.NewUserExistsByIdAdapter(h.grpc)
	service := services.NewActivateEventService(activateEventAdapter, userExistsByIdPort)

	input := services.ActivateEventServiceInput{
		EventId: eventId,
		UserId:  userId,
	}

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
