// Package racecourse
package racecourse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response := gin.H{
		"result": "ok",
	}

	c.IndentedJSON(http.StatusOK, response)
}
