// Package user
package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userModel "tmp-app-backend/model/user"
)

type userService interface {
	GetUsers() ([]userModel.User, error)
}

func Index(service userService) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := service.GetUsers()
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
			return
		}

		c.IndentedJSON(http.StatusOK, response)
	}
}
