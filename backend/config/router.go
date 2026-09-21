// Package config
package config

import (
	raceController "tmp-app-backend/controller/race"
	raceCourseController "tmp-app-backend/controller/race_course"
	userController "tmp-app-backend/controller/user"
	raceService "tmp-app-backend/service/racing"
	userService "tmp-app-backend/service/user"

	"github.com/gin-gonic/gin"
)

func Routing(
	engin *gin.Engine,
	userService *userService.Service,
	raceService *raceService.RaceService,
	raceCourseService *raceService.RaceCourseService,
) {
	// user
	engin.GET("/api/users", userController.Index(userService))
	// race
	engin.GET("/api/races", raceController.Index(raceService))
	// race_course
	engin.GET("/api/race_courses", raceCourseController.Index(raceCourseService))
}
