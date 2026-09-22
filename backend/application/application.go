// Package application
package application

import (
	"database/sql"

	raceInfrastructure "keiba-app-backend/infrastructure/race"
	raceCourseInfrastructure "keiba-app-backend/infrastructure/race_course"
	userInfrastructure "keiba-app-backend/infrastructure/user"
	raceService "keiba-app-backend/service/race"
	raceCourseService "keiba-app-backend/service/race_course"
	raceSearchService "keiba-app-backend/service/race_search"
	userService "keiba-app-backend/service/user"
)

type Application struct {
	UserService       *userService.Service
	RaceService       *raceService.RaceService
	RaceSearchService *raceSearchService.SearchService
	RaceCourseService *raceCourseService.RaceCourseService
}

// NewApplication アプリケーションの依存関係を組み立てる
func NewApplication(db *sql.DB) *Application {
	userStore := userInfrastructure.NewStore(db)
	raceStore := raceInfrastructure.NewStore(db)
	raceCourseStore := raceCourseInfrastructure.NewStore(db)

	return &Application{
		UserService:       userService.NewService(userStore),
		RaceService:       raceService.NewRaceService(raceStore),
		RaceSearchService: raceSearchService.NewSearchService(raceStore),
		RaceCourseService: raceCourseService.NewRaceCourseService(raceCourseStore),
	}
}
