// Package race
package race

import (
	"net/http"

	"github.com/gin-gonic/gin"

	raceService "tmp-app-backend/service/racing"
)

func Index(service *raceService.RaceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := service.GetRaces()
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch races"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}
