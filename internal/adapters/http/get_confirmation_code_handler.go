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

type GetConfirmationCodeHandler struct {
	db   *sql.DB
	grpc *grpc.ClientConn
}

func NewGetConfirmationCodeHandler(db *sql.DB, grpc *grpc.ClientConn) *GetConfirmationCodeHandler {
	return &GetConfirmationCodeHandler{db: db, grpc: grpc}
}

// @Summary Get a confirmation code for an event
// @Description Get a confirmation code for an event
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param userId header int true "User ID"
// @Success 200 {object} services.GetConfirmationCodeOutput
// @Failure 400 {object} services.GetConfirmationCodeOutput
// @Failure 401 {object} services.GetConfirmationCodeOutput
// @Security Bearer
// @Router /confirmation-code/{id} [get]
func (h *GetConfirmationCodeHandler) Handle(c *gin.Context) {
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
	getConfirmationCodePort := database.NewGetConfirmationCodeAdapter(h.db)
	service := services.NewGetConfirmationCodeService(getConfirmationCodePort, userExistsPort)

	input := services.GetConfirmationCodeInput{
		UserId:  userId.(int),
		EventId: eventId,
	}

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
