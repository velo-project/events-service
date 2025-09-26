package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gitlab.com/velo-company/services/events-service/internal/adapters/http"
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

	r := gin.Default()
	r.Use(sentrygin.New(sentrygin.Options{}))

	subscribeHandler := http.NewSubscribeEventHandler(db, grpcConn)
	cancelSubscriptionHandler := http.NewCancelSubscriptionHandler(db)
	confirmSubscriptionHandler := http.NewConfirmSubscriptionHandler(db, grpcConn)
	getConfirmationCodeHandler := http.NewGetConfirmationCodeHandler(db, grpcConn)

	pr := r.Group("/api/events/v1")
	pr.Use(http.AuthMiddleware())
	{
		pr.POST("/subscribe/:id", subscribeHandler.Handle)
		pr.POST("/cancel-subscription/:id", cancelSubscriptionHandler.Handle)
		pr.POST("/confirm-subscription/:id", confirmSubscriptionHandler.Handle)
		pr.GET("/confirmation-code/:id", getConfirmationCodeHandler.Handle)
	}

	r.Run()
}
