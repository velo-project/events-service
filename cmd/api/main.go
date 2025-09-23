package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	r := gin.Default()

	// pr := r.Group("/api/events/v1", http.AuthMiddleware())

	r.Run()
}
