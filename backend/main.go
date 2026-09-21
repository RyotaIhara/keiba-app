package main

import (
	"log"

	"tmp-app-backend/application"
	"tmp-app-backend/config"
	"tmp-app-backend/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := application.New(db)

	engin := gin.Default()
	config.Routing(
		engin,
		app.UserService,
		app.RaceService,
		app.RaceCourseService,
	)
	if err := engin.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
