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
	userStore := userInfrastructure.NewStore(db)
	raceStore := raceInfrastructure.NewStore(db)
	raceCourseStore := raceCourseInfrastructure.NewStore(db)

	return &Application{
		UserService:       userService.NewService(userStore),
		RaceService:       racingService.NewRaceService(raceStore),
		RaceCourseService: racingService.NewRaceCourseService(raceCourseStore),
	}
}
