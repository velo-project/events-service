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

type SubscribeEventHandler struct {
	db   *sql.DB
	grpc *grpc.ClientConn
}

func NewSubscribeEventHandler(db *sql.DB, grpc *grpc.ClientConn) *SubscribeEventHandler {
	return &SubscribeEventHandler{db: db, grpc: grpc}
}

func (h *SubscribeEventHandler) Handle(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "Usuário não está no contexto"})
		return
	}

	eventIdStr := c.Param("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Evento Invalido"})
		return
	}

	userExistsPort := grpcadapter.NewUserExistsByIdAdapter(h.grpc)
	subscribeEventPort := database.NewSubscribeEventAdapter(h.db)
	service := services.NewSubscribeEventService(userExistsPort, subscribeEventPort)

	input := services.SubscribeEventServiceInput{
		UserId:  userId.(int),
		EventId: eventId,
	}

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
