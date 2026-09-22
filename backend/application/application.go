// Package application assembles the application's dependencies.
package application

import (
	"database/sql"

	raceInfrastructure "keiba-app-backend/infrastructure/race"
	raceCourseInfrastructure "keiba-app-backend/infrastructure/race_course"
	userInfrastructure "keiba-app-backend/infrastructure/user"
	raceService "keiba-app-backend/service/race"
	raceCourseService "keiba-app-backend/service/race_course"
	userService "keiba-app-backend/service/user"
)

type Application struct {
	UserService       *userService.Service
	RaceService       *raceService.RaceService
	RaceCourseService *raceCourseService.RaceCourseService
}

func New(db *sql.DB) *Application {
	userStore := userInfrastructure.NewStore(db)
	raceStore := raceInfrastructure.NewStore(db)
	raceCourseStore := raceCourseInfrastructure.NewStore(db)

	return &Application{
		UserService:       userService.NewService(userStore),
		RaceService:       raceService.NewRaceService(raceStore),
		RaceCourseService: raceCourseService.NewRaceCourseService(raceCourseStore),
	}
}
