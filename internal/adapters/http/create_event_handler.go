package http

import (
	"database/sql"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"gitlab.com/velo-company/services/events-service/internal/adapters/ai"
	"gitlab.com/velo-company/services/events-service/internal/adapters/database"
	storageAdapter "gitlab.com/velo-company/services/events-service/internal/adapters/storage"
	"gitlab.com/velo-company/services/events-service/internal/core/services"
)

type CreateEventHandler struct {
	db     *sql.DB
	model  *genai.EmbeddingModel
	bucket *storage.BucketHandle
}

func NewCreateEventHandler(db *sql.DB, md *genai.EmbeddingModel, bucket *storage.BucketHandle) *CreateEventHandler {
	return &CreateEventHandler{
		db:     db,
		model:  md,
		bucket: bucket,
	}
}

func (h *CreateEventHandler) Handle(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{"error": "Image not provided"})
		return
	}

	image, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open image"})
		return
	}
	defer image.Close()

	var input services.CreateEventServiceInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	input.Image = image
	input.ImageExtension = file.Filename[len(file.Filename)-4:]

	createEventPort := database.NewCreateEventAdapter(h.db)
	embeddingsGenerator := ai.NewEmbeddingsGenerator(h.model)
	saveFileAdapter := storageAdapter.NewSaveFileAdapter(h.bucket)
	service := services.NewCreateEventService(createEventPort, embeddingsGenerator, saveFileAdapter)

	output := service.Execute(&input)

	c.JSON(output.StatusCode, output)
}
