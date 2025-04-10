package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	usecases "github.com/velo-project/events-service/assemblers/use_cases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connection to Db
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	postgresServer := os.Getenv("POSTGRES_URI")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPwd := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DATABASE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresServer, postgresUser, postgresPwd, postgresDb, postgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Initializing Api Server
	router := gin.Default()

	createEventController := usecases.CreateEventController(db)

	router.POST("/api/v1/event", createEventController.Perform)

	router.Run(":8080")
}
