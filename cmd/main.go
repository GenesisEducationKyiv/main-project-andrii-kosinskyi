package main

import (
	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/handlers"
	"bitcoin_checker_api/internal/repositories/internal-storage"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	repo, err := internal_storage.NewInternalStorageRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewHandler(cfg, repo)

	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.GET("/rate", handler.Rate)
		v1.POST("/subscription", handler.Subscription)
		v1.POST("/sendEmails", handler.SendEmail)
	}

	router.Run("localhost:8080")
}
