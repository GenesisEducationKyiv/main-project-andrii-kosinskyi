package handler

import (
	"bitcoin_checker_api/internal/validator"
	"fmt"
	"log"
	"net/http"
	"os"

	"bitcoin_checker_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase *usecase.UseCase
}

func NewHandler(u *usecase.UseCase) *Handler {
	return &Handler{
		useCase: u,
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
		c.IndentedJSON(http.StatusBadRequest, ErrInvStatVal)
		return
	}
	log.Printf("endpoint: %s response: %s", handlerEndpoint(c), exchangeRate)
	c.IndentedJSON(http.StatusOK, exchangeRate)
}

func (that *Handler) Subscription(c *gin.Context) {
	log.Printf("endpoint: %s request: %s", handlerEndpoint(c), c.PostForm("email"))
	email := c.PostForm("email")
	if err := validator.ValidMailAddress(email); err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", handlerEndpoint(c), err)
		c.IndentedJSON(http.StatusConflict, ErrInvSubEmail)
		return
	}
	if err := that.useCase.SubscribeEmailOnExchangeRate(email); err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", handlerEndpoint(c), err)
		c.IndentedJSON(http.StatusConflict, ErrInvSubEmail)
		return
	}
	log.Printf("endpoint: %s response: %s", handlerEndpoint(c), EmailAdded)
	c.IndentedJSON(http.StatusOK, EmailAdded)
}

func (that *Handler) SendEmails(c *gin.Context) {
	log.Printf("endpoint: %s request: %s", handlerEndpoint(c), http.NoBody)
	err := that.useCase.SendEmailsWithExchangeRate(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", handlerEndpoint(c), err)
		c.IndentedJSON(http.StatusConflict, ErrEmailsNotSent)
		return
	}
	log.Printf("endpoint: %s response: %s", handlerEndpoint(c), EmailsSent)
	c.IndentedJSON(http.StatusOK, EmailsSent)
}
