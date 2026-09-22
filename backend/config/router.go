// Package config
package config

import (
	raceHandler "tmp-app-backend/handler/race"
	raceCourseHandler "tmp-app-backend/handler/race_course"
	userHandler "tmp-app-backend/handler/user"
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
	engin.GET("/api/users", userHandler.Index(userService))
	// race
	engin.GET("/api/races", raceHandler.Index(raceService))
	engin.POST("/api/races", raceHandler.Create(raceService))
	engin.GET("/api/races/:id", raceHandler.Show(raceService))
	engin.PUT("/api/races/:id", raceHandler.Update(raceService))
	engin.DELETE("/api/races/:id", raceHandler.Delete(raceService))
	// race_course
	engin.GET("/api/race_courses", raceCourseHandler.Index(raceCourseService))
}
