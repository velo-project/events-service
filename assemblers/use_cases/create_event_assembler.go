package usecases

import (
	repositoryImplementation "github.com/velo-project/events-service/adapters/dao/repositories/implementation"
	createevent "github.com/velo-project/events-service/adapters/use_cases/create_event"
	"github.com/velo-project/events-service/controllers"
	usecaseImplementation "github.com/velo-project/events-service/core/use_cases/create_event/implementation"
	"gorm.io/gorm"
)

func CreateEventController(db *gorm.DB) controllers.CreateEventController {
	repository := repositoryImplementation.NewEventRepository(db)
	port := createevent.NewPersistNewEventPortAdapter(repository)
	usecase := usecaseImplementation.NewCreateEventUseCase(port)
	controller := controllers.NewCreateEventController(usecase)

	return controller
}
