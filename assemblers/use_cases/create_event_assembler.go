package usecases

import (
	repositoryImplementation "github.com/velo-project/events-service/adapters/dao/repositories/implementation"
	"github.com/velo-project/events-service/controllers"
	usecaseImplementation "github.com/velo-project/events-service/core/use_cases/create_event/implementation"
	"gorm.io/gorm"
)

func CreateEventController(db *gorm.DB) controllers.CreateEventController {
	repository := repositoryImplementation.NewEventRepository(db)
	usecase := usecaseImplementation.NewCreateEventUseCase(repository)
	controller := controllers.NewCreateEventController(usecase)

	return controller
}
