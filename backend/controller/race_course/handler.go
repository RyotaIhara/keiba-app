// Package racecourse
package racecourse

import (
	"net/http"

	"github.com/gin-gonic/gin"

	raceService "tmp-app-backend/service/racing"
)

func Index(service *raceService.RaceCourseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := service.GetRaceCourses()
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch race courses"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}
