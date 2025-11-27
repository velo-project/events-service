package http

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/getsentry/sentry-go"
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
	getConfirmationCodePort := database.NewGetConfirmationCodeAdapter(h.db)
	service := services.NewGetConfirmationCodeService(getConfirmationCodePort, userExistsPort)

	input := services.GetConfirmationCodeInput{
		UserId:  userId.(int),
		EventId: eventId,
	}

	output, err := service.Execute(&input)
	if err != nil {
		sentry.CaptureException(err)
		msg := gin.H{"message": err.Error()}
		switch err.Error() {
		case "Inscrição não encontrada para este evento":
			c.JSON(http.StatusNotFound, msg)
		case "Este evento já ocorreu":
			c.JSON(http.StatusGone, msg)
		case "Este usuário não existe":
			c.JSON(http.StatusNotFound, msg)
		case "Não foi possível buscar o código de confirmação":
			c.JSON(http.StatusInternalServerError, msg)
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
		}
		return
	}

	c.JSON(http.StatusOK, output)
}
