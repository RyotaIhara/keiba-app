// Package config
package config

import (
	raceController "tmp-app-backend/controller/race"
	raceCourseController "tmp-app-backend/controller/race_course"
	userController "tmp-app-backend/controller/user"

	"github.com/gin-gonic/gin"
)

func Routing(engin *gin.Engine) {
	// user
	engin.GET("/api/users", func(c *gin.Context) {
		userController.Index(c)
	})
	// race
	engin.GET("/api/races", func(c *gin.Context) {
		raceController.Index(c)
	})
	// race_course
	engin.GET("/api/race_courses", func(c *gin.Context) {
		raceCourseController.Index(c)
	})
}
