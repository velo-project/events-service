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
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID de evento inválido"})
		return
	}

	anonymousUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Usuário não autenticado"})
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

	output, err := service.Execute(&input)
	if err != nil {
		sentry.CaptureException(err)
		if errors.Is(err, customErrors.ErrEventNotFound) || err.Error() == "Este usuário não existe" {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
		return
	}

	c.JSON(http.StatusOK, output)
}
