package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	usecases "github.com/velo-project/events-service/assemblers/use_cases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connection to Db
	// TODO: Implement Env Variables to connect to the database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost", "postgres", "1234", "velo_project", "5432",
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
