// Package config
package config

import (
	"keiba-app-backend/application"
	raceHandler "keiba-app-backend/handler/race"
	raceCourseHandler "keiba-app-backend/handler/race_course"
	userHandler "keiba-app-backend/handler/user"

	"github.com/gin-gonic/gin"
)

func Routing(
	engin *gin.Engine,
	app *application.Application,
) {
	// user
	engin.GET("/api/users", userHandler.Index(app.UserService))
	engin.GET("/api/users/:id", userHandler.Show(app.UserService))
	engin.POST("/api/users", userHandler.Create(app.UserService))
	engin.PUT("/api/users/:id", userHandler.Update(app.UserService))
	engin.DELETE("/api/users/:id", userHandler.Delete(app.UserService))
	// race
	engin.GET("/api/races", raceHandler.Index(app.RaceService))
	engin.GET("/api/races/:id", raceHandler.Show(app.RaceService))
	engin.GET("/api/races/:id/details", raceHandler.DetailsByID(app.RaceService))
	engin.POST("/api/races", raceHandler.Create(app.RaceService))
	engin.PUT("/api/races/:id", raceHandler.Update(app.RaceService))
	engin.DELETE("/api/races/:id", raceHandler.Delete(app.RaceService))
	// race_course
	engin.GET("/api/race_courses", raceCourseHandler.Index(app.RaceCourseService))
	engin.GET("/api/race_courses/:id", raceCourseHandler.Show(app.RaceCourseService))
	engin.POST("/api/race_courses", raceCourseHandler.Create(app.RaceCourseService))
	engin.PUT("/api/race_courses/:id", raceCourseHandler.Update(app.RaceCourseService))
	engin.DELETE("/api/race_courses/:id", raceCourseHandler.Delete(app.RaceCourseService))
}
