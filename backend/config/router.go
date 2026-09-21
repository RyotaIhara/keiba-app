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
	engin.POST("/api/races", raceController.Create(raceService))
	engin.GET("/api/races/:id", raceController.Show(raceService))
	engin.PUT("/api/races/:id", raceController.Update(raceService))
	engin.DELETE("/api/races/:id", raceController.Delete(raceService))
	// race_course
	engin.GET("/api/race_courses", raceCourseController.Index(raceCourseService))
}
