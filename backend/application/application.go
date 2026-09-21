// Package application assembles the application's dependencies.
package application

import (
	"database/sql"

	raceInfrastructure "tmp-app-backend/infrastructure/race"
	raceCourseInfrastructure "tmp-app-backend/infrastructure/race_course"
	userInfrastructure "tmp-app-backend/infrastructure/user"
	racingService "tmp-app-backend/service/racing"
	userService "tmp-app-backend/service/user"
)

type Application struct {
	UserService       *userService.Service
	RaceService       *racingService.RaceService
	RaceCourseService *racingService.RaceCourseService
}

func New(db *sql.DB) *Application {
	userRepository := userInfrastructure.NewRepository(db)
	raceRepository := raceInfrastructure.NewRepository(db)
	raceCourseRepository := raceCourseInfrastructure.NewRepository(db)

	return &Application{
		UserService:       userService.NewService(userRepository),
		RaceService:       racingService.NewRaceService(raceRepository),
		RaceCourseService: racingService.NewRaceCourseService(raceCourseRepository),
	}
}
