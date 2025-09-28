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
// @Param code body object{code=string} true "Confirmation code"
// @Success 200 {object} services.ConfirmSubscriptionOutput
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /confirm-subscription/{id} [post]
func (h *ConfirmSubscriptionHandler) Handle(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	eventIdStr := c.Param("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	var requestBody struct {
		Code string `json:"code"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
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

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
