// Package race
package race

import (
	"net/http"

	"github.com/gin-gonic/gin"

	raceService "tmp-app-backend/service/racing"
)

func Index(c *gin.Context) {
	response := raceService.GetRaces()

	c.IndentedJSON(http.StatusOK, response)
}
