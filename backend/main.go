package main

import (
	"log"

	"tmp-app-backend/config"
	"tmp-app-backend/infrastructure/database"
	raceInfrastructure "tmp-app-backend/infrastructure/race"
	raceCourseInfrastructure "tmp-app-backend/infrastructure/race_course"
	userInfrastructure "tmp-app-backend/infrastructure/user"
	raceService "tmp-app-backend/service/racing"
	userService "tmp-app-backend/service/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := userInfrastructure.NewRepository(db)
	raceRepository := raceInfrastructure.NewRepository(db)
	raceCourseRepository := raceCourseInfrastructure.NewRepository(db)

	userSvc := userService.NewService(userRepository)
	raceSvc := raceService.NewRaceService(raceRepository)
	raceCourseSvc := raceService.NewRaceCourseService(raceCourseRepository)

	engin := gin.Default()
	config.Routing(engin, userSvc, raceSvc, raceCourseSvc)
	if err := engin.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
