package handler

import (
	"bitcoin_checker_api/internal/validator"
	"fmt"
	"log"
	"net/http"
	"os"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase *usecase.UseCase
	config  *config.Config
}

func NewHandler(c *config.Config, u *usecase.UseCase) *Handler {
	return &Handler{
		config:  c,
		useCase: u,
	}
}

func (that *Handler) Rate(c *gin.Context) {
	log.Printf("endpoint: %s request: %s", c.Request.Host+c.Request.URL.Path, http.NoBody)
	exchangeRate, err := that.useCase.ExchangeRate(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", c.Request.Host+c.Request.URL.Path, err)
		c.IndentedJSON(http.StatusBadRequest, ErrInvStatVal)
		return
	}
	log.Printf("endpoint: %s response: %s", c.Request.Host+c.Request.URL.Path, exchangeRate)
	c.IndentedJSON(http.StatusOK, exchangeRate)
}

func (that *Handler) Subscription(c *gin.Context) {
	log.Printf("endpoint: %s request: %s", c.Request.Host+c.Request.URL.Path, c.PostForm("email"))
	email := c.PostForm("email")
	if err := validator.ValidMailAddress(email); err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", c.Request.Host+c.Request.URL.Path, err)
		c.IndentedJSON(http.StatusConflict, ErrInvSubEmail)
		return
	}
	if err := that.useCase.SubscribeEmailOnExchangeRate(email); err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", c.Request.Host+c.Request.URL.Path, err)
		c.IndentedJSON(http.StatusConflict, ErrInvSubEmail)
		return
	}
	log.Printf("endpoint: %s response: %s", c.Request.Host+c.Request.URL.Path, EmailAdded)
	c.IndentedJSON(http.StatusOK, EmailAdded)
}

func (that *Handler) SendEmails(c *gin.Context) {
	log.Printf("endpoint: %s request: %s", c.Request.Host+c.Request.URL.Path, http.NoBody)
	err := that.useCase.SendEmailsWithExchangeRate(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "endpoint: %s error: %s", c.Request.Host+c.Request.URL.Path, err)
		c.IndentedJSON(http.StatusConflict, ErrEmailsNotSent)
		return
	}
	log.Printf("endpoint: %s response: %s", c.Request.Host+c.Request.URL.Path, EmailsSent)
	c.IndentedJSON(http.StatusOK, EmailsSent)
}
