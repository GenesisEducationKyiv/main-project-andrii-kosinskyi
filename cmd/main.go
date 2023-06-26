package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"bitcoin_checker_api/internal/usecase"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/handler"
	"bitcoin_checker_api/internal/repository/storage"

	"github.com/gin-gonic/gin"
)

//nolint:all
func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt(cfg.Server.Port, 10),
		Handler: initApp(cfg),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func initApp(cfg *config.Config) http.Handler {
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

	return router
}
