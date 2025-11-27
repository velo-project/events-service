package http

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/velo-company/services/events-service/internal/adapters/database"
	grpcadapter "gitlab.com/velo-company/services/events-service/internal/adapters/grpc"
	customErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID de evento inválido"})
		return
	}

	anonymousUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Usuário não autenticado"})
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

	output, err := service.Execute(&input)
	if err != nil {
		sentry.CaptureException(err)
		if errors.Is(err, customErrors.ErrBlockedActivateEvent) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, customErrors.ErrEventNotFound) || err.Error() == "Este usuário não existe" {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
		return
	}

	c.JSON(http.StatusOK, output)
}
