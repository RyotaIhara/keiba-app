package main

import (
	"tmp-app-backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	engin := gin.Default()
	config.Routing(engin)
	engin.Run(":3000")
}
