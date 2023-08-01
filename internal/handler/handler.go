package handler

import (
	"fmt"
	"net/http"

	"bitcoin_checker_api/internal/logger"
	"bitcoin_checker_api/internal/renderer"

	"bitcoin_checker_api/internal/validator"

	"bitcoin_checker_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase       *usecase.UseCase
	renderService *renderer.Render
	log           logger.ILog
}

func NewHandler(u *usecase.UseCase, r *renderer.Render, l logger.ILog) *Handler {
	return &Handler{
		useCase:       u,
		renderService: r,
		log:           l,
	}
}

func handlerEndpoint(c *gin.Context) string {
	return fmt.Sprintf("%s%s", c.Request.Host, c.Request.URL.Path)
}

func (that *Handler) Rate(c *gin.Context) {
	that.log.Info(c, fmt.Sprintf("endpoint: %s request: %s", handlerEndpoint(c), http.NoBody))
	exchangeRate, err := that.useCase.ExchangeRate(c)
	if err != nil {
		that.log.Error(c, fmt.Sprintf("endpoint: %s error: %s", handlerEndpoint(c), err))
		that.renderService.SendError(c, http.StatusBadRequest, ErrInvStatVal)
		return
	}
	that.log.Info(c, fmt.Sprintf("endpoint: %s response: %v", handlerEndpoint(c), exchangeRate))
	that.renderService.SendSuccess(c, exchangeRate)
}

func (that *Handler) Subscription(c *gin.Context) {
	that.log.Info(c, fmt.Sprintf("endpoint: %s request: %s", handlerEndpoint(c), c.PostForm("email")))
	email := c.PostForm("email")
	if err := validator.ValidMailAddress(email); err != nil {
		that.log.Error(c, fmt.Sprintf("endpoint: %s error: %s", handlerEndpoint(c), err))
		that.renderService.SendError(c, http.StatusConflict, ErrInvSubEmail)
		return
	}
	if err := that.useCase.SubscribeEmailOnExchangeRate(email); err != nil {
		that.log.Error(c, fmt.Sprintf("endpoint: %s error: %s", handlerEndpoint(c), err))
		that.renderService.SendError(c, http.StatusConflict, ErrInvSubEmail)
		return
	}
	that.log.Info(c, fmt.Sprintf("endpoint: %s response: %s", handlerEndpoint(c), EmailAdded))
	that.renderService.SendSuccess(c, EmailAdded)
}

func (that *Handler) SendEmails(c *gin.Context) {
	that.log.Info(c, fmt.Sprintf("endpoint: %s request: %s", handlerEndpoint(c), http.NoBody))
	err := that.useCase.SendEmailsWithExchangeRate(c)
	if err != nil {
		that.log.Error(c, fmt.Sprintf("endpoint: %s error: %s", handlerEndpoint(c), err))
		that.renderService.SendError(c, http.StatusConflict, ErrEmailsNotSent)
		return
	}
	that.log.Info(c, fmt.Sprintf("endpoint: %s response: %s", handlerEndpoint(c), EmailsSent))
	that.renderService.SendSuccess(c, EmailsSent)
}
