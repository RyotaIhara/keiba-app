// Package config
package config

import (
	indeController "tmp-app-backend/controller/index"

	"github.com/gin-gonic/gin"
)

func Routing(engin *gin.Engine) {
	// index
	engin.GET("/api/index", func(c *gin.Context) {
		indeController.Index(c)
	})
}
