// Package application assembles the application's dependencies.
package application

import (
	"database/sql"

	raceInfrastructure "keiba-app-backend/infrastructure/race"
	raceCourseInfrastructure "keiba-app-backend/infrastructure/race_course"
	userInfrastructure "keiba-app-backend/infrastructure/user"
	racingService "keiba-app-backend/service/racing"
	userService "keiba-app-backend/service/user"
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
