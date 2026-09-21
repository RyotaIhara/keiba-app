// Package racecourse
package racecourse

import (
	"net/http"

	"github.com/gin-gonic/gin"

	raceService "tmp-app-backend/service/racing"
)

func Index(c *gin.Context) {
	response := raceService.GetRaceCourses()

	c.IndentedJSON(http.StatusOK, response)
}
