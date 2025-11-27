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

type SubscribeEventHandler struct {
	db   *sql.DB
	grpc *grpc.ClientConn
}

func NewSubscribeEventHandler(db *sql.DB, grpc *grpc.ClientConn) *SubscribeEventHandler {
	return &SubscribeEventHandler{db: db, grpc: grpc}
}

// @Summary Subscribe to an event
// @Description Subscribe to an event
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param userId header int true "User ID"
// @Success 200 {object} services.SubscribeEventServiceOutput
// @Failure 400 {object} services.SubscribeEventServiceOutput
// @Failure 401 {object} services.SubscribeEventServiceOutput
// @Security Bearer
// @Router /subscribe/{id} [post]
func (h *SubscribeEventHandler) Handle(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"message": "Usuário não está no contexto"})
		return
	}

	eventIdStr := c.Param("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "ID de evento inválido"})
		return
	}

	userExistsPort := grpcadapter.NewUserExistsByIdAdapter(h.grpc)
	subscribeEventPort := database.NewSubscribeEventAdapter(h.db)
	service := services.NewSubscribeEventService(userExistsPort, subscribeEventPort)

	input := services.SubscribeEventServiceInput{
		UserId:  userId.(int),
		EventId: eventId,
	}

	output, err := service.Execute(&input)
	if err != nil {
		sentry.CaptureException(err)
		if errors.Is(err, customErrors.ErrUserAlreadySubscribed) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if errors.Is(err, customErrors.ErrEventNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		if err.Error() == "Este usuário não existe" {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
		return
	}

	c.JSON(http.StatusCreated, output)
}
