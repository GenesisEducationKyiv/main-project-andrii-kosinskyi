package main

import (
	"log"
	"strconv"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/handlers"
	"bitcoin_checker_api/internal/repositories/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	repo, err := storage.NewInternalStorageRepository(cfg)
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

	err = router.Run(":" + strconv.FormatInt(cfg.Service.Port, 10))
	if err != nil {
		return
	}
}
