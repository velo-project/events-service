package http

import (
	"database/sql"
	"errors"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"gitlab.com/velo-company/services/events-service/internal/adapters/ai"
	"gitlab.com/velo-company/services/events-service/internal/adapters/database"
	grpcadapter "gitlab.com/velo-company/services/events-service/internal/adapters/grpc"
	storageAdapter "gitlab.com/velo-company/services/events-service/internal/adapters/storage"
	customErrors "gitlab.com/velo-company/services/events-service/internal/core/errors"
	"gitlab.com/velo-company/services/events-service/internal/core/services"
	"google.golang.org/grpc"
)

type CreateEventHandler struct {
	db     *sql.DB
	model  *genai.EmbeddingModel
	bucket *storage.BucketHandle
	grpc   *grpc.ClientConn
}

func NewCreateEventHandler(db *sql.DB, md *genai.EmbeddingModel, bucket *storage.BucketHandle, grpcConn *grpc.ClientConn) *CreateEventHandler {
	return &CreateEventHandler{
		db:     db,
		model:  md,
		bucket: bucket,
		grpc:   grpcConn,
	}
}

// @Summary Create an event
// @Description Create an event
// @Tags events
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Event image"
// @Param name formData string true "Event name"
// @Param description formData string false "Event description"
// @Param date formData string true "Event date"
// @Param location formData string false "Event location"
// @Success 201 {object} services.CreateEventServiceOutput
// @Failure 400 {object} services.CreateEventServiceOutput
// @Failure 500 {object} services.CreateEventServiceOutput
// @Security Bearer
// @Router /create [post]
func (h *CreateEventHandler) Handle(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"message": "Imagem não fornecida", "status_code": 400})
		return
	}

	image, err := file.Open()
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(500, gin.H{"message": "Falha ao abrir a imagem", "status_code": 500})
		return
	}
	defer image.Close()

	anonymousUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(401, gin.H{"message": "Usuário não autenticado", "status_code": 401})
		return
	}

	userId := anonymousUserId.(int)

	var input services.CreateEventServiceInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"message": "Entrada inválida", "status_code": 400})
		return
	}
	input.Image = image
	input.ImageExtension = file.Filename[len(file.Filename)-4:]
	input.UserId = userId

	createEventPort := database.NewCreateEventAdapter(h.db)
	embeddingsGenerator := ai.NewEmbeddingsGenerator(h.model)
	saveFileAdapter := storageAdapter.NewSaveFileAdapter(h.bucket)
	userExistsByIdPort := grpcadapter.NewUserExistsByIdAdapter(h.grpc)
	service := services.NewCreateEventService(createEventPort, embeddingsGenerator, saveFileAdapter, userExistsByIdPort)

	output, err := service.Execute(&input)
	if err != nil {
		sentry.CaptureException(err)
		if errors.Is(err, customErrors.ErrEventNotCreated) {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		if err.Error() == "Este usuário não existe" {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		if err.Error() == "Extensão de imagem inválida" {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Ocorreu um erro interno"})
		return
	}

	c.JSON(http.StatusCreated, output)
}
