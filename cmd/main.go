//nolint:gocritic,gosec,gomnd,gosimple,govet,staticcheck,goimports // issues only in graceful shutdown process
package main

import (
	"bitcoin_checker_api/internal"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"bitcoin_checker_api/config"
)

func main() {
	cfg := config.NewConfig()
	if err := cfg.Load(); err != nil {
		log.Fatalf("cfg.Load: %s\n", err)
	}

	h, err := internal.InitApp(cfg)
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
