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

type ConfirmSubscriptionHandler struct {
	db   *sql.DB
	grpc *grpc.ClientConn
}

func NewConfirmSubscriptionHandler(db *sql.DB, grpc *grpc.ClientConn) *ConfirmSubscriptionHandler {
	return &ConfirmSubscriptionHandler{db: db, grpc: grpc}
}

// @Summary Confirm a subscription to an event
// @Description Confirm a subscription to an event
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param userId header int true "User ID"
// @Param code body object{code=string} true "Confirmation code"
// @Success 200 {object} services.ConfirmSubscriptionOutput
// @Failure 400 {object} services.ConfirmSubscriptionOutput
// @Failure 401 {object} services.ConfirmSubscriptionOutput
// @Security Bearer
// @Router /confirm-subscription/{id} [post]
func (h *ConfirmSubscriptionHandler) Handle(c *gin.Context) {
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

	var requestBody struct {
		Code string `json:"code"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Corpo da requisição inválido"})
		return
	}

	userExistsPort := grpcadapter.NewUserExistsByIdAdapter(h.grpc)
	confirmSubscriptionPort := database.NewConfirmSubscriptionAdapter(h.db)
	service := services.NewConfirmSubscriptionService(confirmSubscriptionPort, userExistsPort)

	input := services.ConfirmSubscriptionInput{
		Code:    requestBody.Code,
		UserId:  userId.(int),
		EventId: eventId,
	}

	output, err := service.Execute(&input)
	if err != nil {
		sentry.CaptureException(err)
		if errors.Is(err, customErrors.ErrInvalidConfirmationCode) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err.Error() == "Este usuário não existe" {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
		return
	}

	c.JSON(http.StatusOK, output)
}
