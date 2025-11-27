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

type CancelSubscriptionHandler struct {
	db   *sql.DB
	grpc *grpc.ClientConn
}

func NewCancelSubscriptionHandler(db *sql.DB, grpc *grpc.ClientConn) *CancelSubscriptionHandler {
	return &CancelSubscriptionHandler{db: db, grpc: grpc}
}

// @Summary Cancel a subscription to an event
// @Description Cancel a subscription to an event
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param userId header int true "User ID"
// @Success 200 {object} services.CancelSubscriptionServiceOutput
// @Failure 400 {object} services.CancelSubscriptionServiceOutput
// @Failure 401 {object} services.CancelSubscriptionServiceOutput
// @Security Bearer
// @Router /cancel-subscription/{id} [post]
func (h *CancelSubscriptionHandler) Handle(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Usuário não autenticado"})
		return
	}

	eventIdStr := c.Param("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID de evento inválido"})
		return
	}

	userExistsPort := grpcadapter.NewUserExistsByIdAdapter(h.grpc)
	cancelSubscriptionPort := database.NewCancelSubscriptionAdapter(h.db)
	service := services.NewCancelSubscriptionService(cancelSubscriptionPort, userExistsPort)

	input := services.CancelSubscriptionServiceInput{
		UserId:  userId.(int),
		EventId: eventId,
	}

	output, err := service.Execute(&input)
	if err != nil {
		sentry.CaptureException(err)
		if errors.Is(err, customErrors.ErrBlockedCancelSubscription) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, customErrors.ErrEventNotFound) || errors.Is(err, customErrors.ErrUserSubscriptionNotFound) || err.Error() == "Este usuário não existe" {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
		return
	}

	c.JSON(http.StatusOK, output)
}
