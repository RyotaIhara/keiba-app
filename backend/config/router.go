// Package config
package config

import (
	raceHandler "keiba-app-backend/handler/race"
	raceCourseHandler "keiba-app-backend/handler/race_course"
	userHandler "keiba-app-backend/handler/user"
	raceService "keiba-app-backend/service/race"
	raceCourseService "keiba-app-backend/service/race_course"
	userService "keiba-app-backend/service/user"

	"github.com/gin-gonic/gin"
)

func Routing(
	engin *gin.Engine,
	userService *userService.Service,
	raceService *raceService.RaceService,
	raceCourseService *raceCourseService.RaceCourseService,
) {
	// user
	engin.GET("/api/users", userHandler.Index(userService))
	engin.GET("/api/users/:id", userHandler.Show(userService))
	engin.POST("/api/users", userHandler.Create(userService))
	engin.PUT("/api/users/:id", userHandler.Update(userService))
	engin.DELETE("/api/users/:id", userHandler.Delete(userService))
	// race
	engin.GET("/api/races", raceHandler.Index(raceService))
	engin.POST("/api/races", raceHandler.Create(raceService))
	engin.GET("/api/races/:id/details", raceHandler.DetailsIndex(raceService))
	engin.GET("/api/races/:id/details/:race_detail_id", raceHandler.DetailsShow(raceService))
	engin.GET("/api/races/:id", raceHandler.Show(raceService))
	engin.PUT("/api/races/:id", raceHandler.Update(raceService))
	engin.DELETE("/api/races/:id", raceHandler.Delete(raceService))
	// race_course
	engin.GET("/api/race_courses", raceCourseHandler.Index(raceCourseService))
	engin.GET("/api/race_courses/:id", raceCourseHandler.Show(raceCourseService))
	engin.POST("/api/race_courses", raceCourseHandler.Create(raceCourseService))
	engin.PUT("/api/race_courses/:id", raceCourseHandler.Update(raceCourseService))
	engin.DELETE("/api/race_courses/:id", raceCourseHandler.Delete(raceCourseService))
}
