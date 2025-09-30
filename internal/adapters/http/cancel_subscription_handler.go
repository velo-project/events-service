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

type CancelSubscriptionHandler struct {
	db   *sql.DB
	grpc *grpc.ClientConn
}

func NewCancelSubscriptionHandler(db *sql.DB) *CancelSubscriptionHandler {

	return &CancelSubscriptionHandler{db: db}
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
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	eventIdStr := c.Param("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	userExistsPort := grpcadapter.NewUserExistsByIdAdapter(h.grpc)
	cancelSubscriptionPort := database.NewCancelSubscriptionAdapter(h.db)
	service := services.NewCancelSubscriptionService(cancelSubscriptionPort, userExistsPort)

	input := services.CancelSubscriptionServiceInput{
		UserId:  userId.(int),
		EventId: eventId,
	}

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
