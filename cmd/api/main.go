package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gitlab.com/velo-company/services/events-service/docs"
	"gitlab.com/velo-company/services/events-service/internal/adapters/database"
	"gitlab.com/velo-company/services/events-service/internal/adapters/http"
	"gitlab.com/velo-company/services/events-service/internal/core/services"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Events Service API
// @version 1.0
// @description This is the API for the Events Service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/events/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.BasePath = "/api/events/v1"
	if err := godotenv.Load(); err != nil {
		log.Print("WARN: No .env file, using default system variables")
	}

	postgresConn := os.Getenv("POSTGRES_CONNECTION_STRING")

	db, err := sql.Open("postgres", postgresConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	grpcStr := os.Getenv("USER_SERVICE_GRPC_ADDRESS")
	grpcConn, err := grpc.NewClient(grpcStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Can't connect to User Service gRPC Server: %v", err)
	}
	defer grpcConn.Close()

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	geminiApiKey := os.Getenv("GEMINI_API_KEY")
	geminiModel := os.Getenv("GEMINI_EMBEDDINGS_MODEL")

	if geminiModel == "" {
		panic("GEMINI_EMBEDDINGS_MODEL environment variable not set")
	}
	if geminiApiKey == "" {
		panic("GEMINI_API_KEY environment variable not set")
	}

	ctx := context.Background()
	geminiClient, err := genai.NewClient(ctx, option.WithAPIKey(geminiApiKey))

	if err != nil {
		panic(err)
	}

	defer geminiClient.Close()

	embeddingsModel := geminiClient.EmbeddingModel(geminiModel)

	gcsClient, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GCP_APPLICATION_CREDENTIALS")))

	if err != nil {
		panic(err)
	}
	defer gcsClient.Close()

	bucketName := os.Getenv("GCP_BUCKET_NAME")

	bucket := gcsClient.Bucket(bucketName)

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(sentrygin.New(sentrygin.Options{}))

	// Adapters
	getRecommendedEventsAdapter := database.NewGetRecommendedEventsAdapter(db)
	getTrendingEventsAdapter := database.NewGetTrendingEventsAdapter(db)
	getLastParticipatedEventsAdapter := database.NewGetLastParticipatedEventsAdapter(db)
	getSubscribedEventsAdapter := database.NewGetSubscribedEventsAdapter(db)

	// Services
	getRecommendedEventsService := services.NewGetRecommendedEventsService(getRecommendedEventsAdapter)
	getTrendingEventsService := services.NewGetTrendingEventsService(getTrendingEventsAdapter)
	getLastParticipatedEventsService := services.NewGetLastParticipatedEventsService(getLastParticipatedEventsAdapter)
	getSubscribedEventsService := services.NewGetSubscribedEventsService(getSubscribedEventsAdapter)

	subscribeHandler := http.NewSubscribeEventHandler(db, grpcConn)
	cancelSubscriptionHandler := http.NewCancelSubscriptionHandler(db, grpcConn)
	confirmSubscriptionHandler := http.NewConfirmSubscriptionHandler(db, grpcConn)
	getConfirmationCodeHandler := http.NewGetConfirmationCodeHandler(db, grpcConn)
	createEventHandler := http.NewCreateEventHandler(db, embeddingsModel, bucket, grpcConn)
	cancelEventHandler := http.NewCancelEventHandler(db, grpcConn)
	suspendEventHandler := http.NewSuspendEventHandler(db, grpcConn)
	activateEventHandler := http.NewActivateEventHandler(db, grpcConn)
	getEventsHandler := http.NewGetEventsHandler(getRecommendedEventsService, getTrendingEventsService, getLastParticipatedEventsService, getSubscribedEventsService)

	pr := r.Group("/api/events/v1")
	pr.Use(http.AuthMiddleware([]string{"USER", "ADMIN"}))
	{
		pr.POST("/subscribe/:id", subscribeHandler.Handle)
		pr.POST("/cancel-subscription/:id", cancelSubscriptionHandler.Handle)
		pr.POST("/confirm-subscription/:id", confirmSubscriptionHandler.Handle)
		pr.GET("/confirmation-code/:id", getConfirmationCodeHandler.Handle)
		pr.POST("/create", http.AuthMiddleware([]string{"ENTERPRISE", "ADMIN"}), createEventHandler.Handle)
		pr.PATCH("/cancel/:id", http.AuthMiddleware([]string{"ENTERPRISE", "ADMIN"}), cancelEventHandler.Handle)
		pr.PATCH("/suspend/:id", http.AuthMiddleware([]string{"ENTERPRISE", "ADMIN"}), suspendEventHandler.Handle)
		pr.PATCH("/activate/:id", http.AuthMiddleware([]string{"ENTERPRISE", "ADMIN"}), activateEventHandler.Handle)
		pr.GET("/events", getEventsHandler.Handle)
	}

	pr.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
