package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bitcoin_checker_api/internal/renderer"

	"bitcoin_checker_api/internal/validator"

	"bitcoin_checker_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase       *usecase.UseCase
	renderService *renderer.Render
}

func NewHandler(u *usecase.UseCase, r *renderer.Render) *Handler {
	return &Handler{
		useCase:       u,
		renderService: r,
	}
}

func handlerEndpoint(c *gin.Context) string {
	return fmt.Sprintf("%s%s", c.Request.Host, c.Request.URL.Path)
}

func (that *Handler) Rate(c *gin.Context) {
	log.Printf("endpoint: %s request: %s", handlerEndpoint(c), http.NoBody)
	exchangeRate, err := that.useCase.ExchangeRate(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", handlerEndpoint(c), err)
		that.renderService.SendError(c, http.StatusBadRequest, ErrInvStatVal)
		return
	}
	log.Printf("endpoint: %s response: %v", handlerEndpoint(c), exchangeRate)
	that.renderService.SendSuccess(c, exchangeRate)
}

func (that *Handler) Subscription(c *gin.Context) {
	log.Printf("endpoint: %s request: %s", handlerEndpoint(c), c.PostForm("email"))
	email := c.PostForm("email")
	if err := validator.ValidMailAddress(email); err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", handlerEndpoint(c), err)
		that.renderService.SendError(c, http.StatusConflict, ErrInvSubEmail)
		return
	}
	if err := that.useCase.SubscribeEmailOnExchangeRate(email); err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", handlerEndpoint(c), err)
		that.renderService.SendError(c, http.StatusConflict, ErrInvSubEmail)
		return
	}
	log.Printf("endpoint: %s response: %s", handlerEndpoint(c), EmailAdded)
	that.renderService.SendSuccess(c, EmailAdded)
}

func (that *Handler) SendEmails(c *gin.Context) {
	log.Printf("endpoint: %s request: %s", handlerEndpoint(c), http.NoBody)
	err := that.useCase.SendEmailsWithExchangeRate(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", handlerEndpoint(c), err)
		that.renderService.SendError(c, http.StatusConflict, ErrEmailsNotSent)
		return
	}
	log.Printf("endpoint: %s response: %s", handlerEndpoint(c), EmailsSent)
	that.renderService.SendSuccess(c, EmailsSent)
}
