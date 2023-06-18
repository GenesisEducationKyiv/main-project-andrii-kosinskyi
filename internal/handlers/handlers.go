package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/smtp"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/models"
	"bitcoin_checker_api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository repositories.Repository
	cfg        *config.Config
}

func NewHandler(cfg *config.Config, repository repositories.Repository) *Handler {
	return &Handler{
		cfg:        cfg,
		repository: repository,
	}
}

func rate(cfg *config.Config) (string, error) {
	ctx := context.Background()
	converter := models.NewConverter()
	requestURL := fmt.Sprintf("%s%s", cfg.Converter.Endpoint, converter.GetQueryParams())
	res, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return "", err
	}
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()
	return string(body), nil
}

func sendMail(email, data string) {
	from := "from@gmail.com"
	password := "<Email Password>"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(data)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Printf("Error: %s", err) //nolint:forbidigo // log for now
		return
	}
}

func (that *Handler) Rate(c *gin.Context) {
	data, err := rate(that.cfg)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid status value")
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (that *Handler) Subscription(c *gin.Context) {
	email := c.PostForm("email")
	err := that.repository.Write(email)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, email)
		return
	}
	c.IndentedJSON(http.StatusOK, "E-mail додано")
}

func (that *Handler) SendEmail(c *gin.Context) {
	users := that.repository.ReadAll()
	data, _ := rate(that.cfg)
	for _, user := range users {
		sendMail(user.Email, data)
	}
	c.IndentedJSON(http.StatusOK, "E-mailʼи відправлено")
}
