package usecases

import (
	"github.com/velo-project/events-service/controllers"
	"github.com/velo-project/events-service/core/use_cases/create_event/implementation"
	"gorm.io/gorm"
)

func CreateEventController(db *gorm.DB) controllers.CreateEventController {
	usecase := implementation.NewCreateEventUseCase()
	controller := controllers.NewCreateEventController(&usecase)

	return controller
}
