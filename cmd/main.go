package main

import (
	"github.com/gin-gonic/gin"
	usecases "github.com/velo-project/events-service/assemblers/use_cases"
)

func main() {
	router := gin.Default()

	createEventController := usecases.CreateEventController()

	router.POST("/api/v1/event", createEventController.Perform)

	router.Run(":8080")
}
