package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"keiba-app-backend/application"
	"keiba-app-backend/config"
	"keiba-app-backend/infrastructure/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// main アプリケーションを起動する
func main() {
	origins, err := allowedOrigins()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := application.New(db)

	engin := gin.Default()
	engin.Use(cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))
	config.Routing(
		engin,
		app,
	)
	if err := engin.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}

// allowedOrigins 環境変数から許可するオリジンを取得する
func allowedOrigins() ([]string, error) {
	origins := os.Getenv("CORS_ALLOW_ORIGINS")
	if origins == "" {
		return nil, fmt.Errorf("CORS_ALLOW_ORIGINS is required")
	}

	allowed := make([]string, 0)
	for origin := range strings.SplitSeq(origins, ",") {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowed = append(allowed, origin)
		}
	}

	if len(allowed) == 0 {
		return nil, fmt.Errorf("CORS_ALLOW_ORIGINS must contain at least one origin")
	}

	return allowed, nil
}
