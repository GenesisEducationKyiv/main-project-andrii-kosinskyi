package main

import (
	"fmt"
	"log"
	"strconv"

	"bitcoin_checker_api/internal/usecase"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/handler"
	"bitcoin_checker_api/internal/repository/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	repo, err := storage.NewStorageRepository(&cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}

	useCaseConfig := &usecase.UseCaseConfig{
		ExchangeRate: &cfg.ExchangeRate,
		EmailService: &cfg.EmailService,
	}

	h := handler.NewHandler(cfg, usecase.NewUseCase(useCaseConfig, repo))

	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.GET("/rate", h.Rate)
		v1.POST("/subscribe", h.Subscription)
		v1.POST("/sendEmails", h.SendEmails)
	}

	fmt.Printf(strconv.FormatInt(cfg.Server.Port, 10))
	err = router.Run(":" + strconv.FormatInt(cfg.Server.Port, 10))
	if err != nil {
		return
	}
}
