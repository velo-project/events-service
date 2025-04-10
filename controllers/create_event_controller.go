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
	var input io.CreateEventInput
	ctx.ShouldBindBodyWithJSON(&input)

	result := cec.usecase.Execute(input)

	ctx.JSON(result.StatusCode, result)
}
