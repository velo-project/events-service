package controllers

import (
	"github.com/gin-gonic/gin"
	createevent "github.com/velo-project/events-service/core/use_cases/create_event"
	"github.com/velo-project/events-service/core/use_cases/create_event/io"
)

type CreateEventController struct {
	usecase createevent.CreateEventUseCase
}

func NewCreateEventController(usecase createevent.CreateEventUseCase) CreateEventController {
	return CreateEventController{
		usecase: usecase,
	}
}

func (cec *CreateEventController) Perform(ctx *gin.Context) {
	input := io.CreateEventInput{}
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message":    "Invalid input",
			"error":      err.Error(),
			"statusCode": 400,
		})
		return
	}

	if errs := input.Validate(); errs != nil {
		ctx.JSON(400, gin.H{
			"message":    "Invalid input",
			"errors":     errs,
			"statusCode": 400,
		})
		return
	}

	result := cec.usecase.Execute(input)

	ctx.JSON(result.StatusCode, result)
}
