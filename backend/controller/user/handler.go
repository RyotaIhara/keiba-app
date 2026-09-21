// Package user
package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userService "tmp-app-backend/service/user"
)

func Index(c *gin.Context) {
	response := userService.GetUsers()

	c.IndentedJSON(http.StatusOK, response)
}
