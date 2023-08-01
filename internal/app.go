package internal

import (
	"log"
	"net/http"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/handler"
	"bitcoin_checker_api/internal/logger"
	"bitcoin_checker_api/internal/pkg/broker"
	"bitcoin_checker_api/internal/pkg/email"
	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/renderer"
	"bitcoin_checker_api/internal/repository"
	"bitcoin_checker_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

func InitApp(cfg *config.Config) (http.Handler, error) {
	h, err := initServices(cfg)
	if err != nil {
		return nil, err
	}

	return initRouter(h)
}

func initRouter(h *handler.Handler) (http.Handler, error) {
	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.GET("/rate", h.Rate)
		v1.POST("/subscribe", h.Subscription)
		v1.POST("/sendEmails", h.SendEmails)
	}

	return router, nil
}

func initServices(cfg *config.Config) (*handler.Handler, error) {
	repoServ, err := repository.NewLocalRepository(&cfg.Storage)
	if err != nil {
		return nil, err
	}

	excRateServ, err := initRateChain(cfg.ExchangeRate)
	if err != nil {
		return nil, err
	}

	rabbitmqConn, err := broker.NewRabbitMQ(cfg.RabbitMQ)
	if err != nil {
		return nil, err
	}

	kafkaConn, err := broker.NewRabbitMQ(cfg.RabbitMQ)
	if err != nil {
		return nil, err
	}

	log.Printf("\ncfg.RabbitMQ %#v\n", cfg.RabbitMQ)

	brokerMap := map[string]broker.Connection{
		broker.RabbitmqBrokerName: rabbitmqConn,
		broker.KafkaBrokerName:    kafkaConn,
	}
	logBroker, err := broker.LoggerServiceFactory(cfg.Logger.Broker, brokerMap)
	if err != nil {
		return nil, err
	}
	logServ := logger.NewLog(logBroker)

	emailServ := email.NewService(&cfg.EmailService)
	renderServ := renderer.NewRender()
	useCase := usecase.NewUseCase(repoServ, excRateServ, emailServ)

	return handler.NewHandler(useCase, renderServ, logServ), nil
}

func initRateChain(cfg config.ExchangeRate) (exchangerate.ExchangeRater, error) {
	excRate, err := exchangerate.NewExchangeRate(cfg.Coinpaprika)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	excRateBinance, err := exchangerate.NewExchangeRate(cfg.Binance)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	excRate.SetNext(excRateBinance)
	return excRate, nil
}
