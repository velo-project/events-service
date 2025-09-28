package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gitlab.com/velo-company/services/events-service/internal/adapters/http"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
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
	geminiModel := os.Getenv("GEMINI_MODEL")

	if geminiModel == "" {
		panic("GEMINI_MODEL environment variable not set")
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

	r := gin.Default()
	r.Use(sentrygin.New(sentrygin.Options{}))

	subscribeHandler := http.NewSubscribeEventHandler(db, grpcConn)
	cancelSubscriptionHandler := http.NewCancelSubscriptionHandler(db)
	confirmSubscriptionHandler := http.NewConfirmSubscriptionHandler(db, grpcConn)
	getConfirmationCodeHandler := http.NewGetConfirmationCodeHandler(db, grpcConn)
	createEventHandler := http.NewCreateEventHandler(db, embeddingsModel)

	pr := r.Group("/api/events/v1")
	pr.Use(http.AuthMiddleware([]string{"USER", "ADMIN"}))
	{
		pr.POST("/subscribe/:id", subscribeHandler.Handle)
		pr.POST("/cancel-subscription/:id", cancelSubscriptionHandler.Handle)
		pr.POST("/confirm-subscription/:id", confirmSubscriptionHandler.Handle)
		pr.GET("/confirmation-code/:id", getConfirmationCodeHandler.Handle)
		pr.POST("/create", http.AuthMiddleware([]string{"ENTERPRISE", "ADMIN"}), createEventHandler.Handle)
	}

	r.Run()
}
