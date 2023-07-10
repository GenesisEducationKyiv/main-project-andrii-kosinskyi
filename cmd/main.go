//nolint:gocritic,gosec,gomnd,gosimple,govet,staticcheck,goimports // issues only in graceful shutdown process
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

	"bitcoin_checker_api/internal/pkg/email"
	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/repository"

	"bitcoin_checker_api/internal/usecase"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatalf("cfg.Load: %s\n", err)
	}

	h, err := initApp(cfg)
	if err != nil {
		log.Fatalf("initApp: %s\n", err)
	}
	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt(cfg.Server.Port, 10),
		Handler: h,
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

func initApp(cfg *config.Config) (http.Handler, error) {
	repo, err := repository.NewLocalRepository(&cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}

	excRate, err := exchangerate.NewExchangeRate(cfg.ExchangeRate.Coinpaprika)
	if err != nil {
		log.Fatal(err)
	}
	excRateBinance, err := exchangerate.NewExchangeRate(cfg.ExchangeRate.Binance)
	if err != nil {
		log.Fatal(err)
	}
	excRate.SetNext(excRateBinance)
	emailServ := email.NewService(&cfg.EmailService)

	h := handler.NewHandler(usecase.NewUseCase(repo, excRate, emailServ))

	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.GET("/rate", h.Rate)
		v1.POST("/subscribe", h.Subscription)
		v1.POST("/sendEmails", h.SendEmails)
	}

	return router, nil
}
