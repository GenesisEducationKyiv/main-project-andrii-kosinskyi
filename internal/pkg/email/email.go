package email

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bitcoin_checker_api/config"

	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/sendgrid/sendgrid-go"
)

type Service struct {
	APIKey      string
	FromAddress string
	FromName    string
}

func NewService(c *config.EmailService) *Service {
	return &Service{
		APIKey:      c.APIKey,
		FromAddress: c.FromAddress,
		FromName:    c.FromName,
	}
}

func (that *Service) Send(email, data string) error {
	from := mail.NewEmail(that.FromName, that.FromAddress)
	subject := "Current exchange rate by your subscription"
	to := mail.NewEmail("Dear customer", email)
	plainTextContent := "Current exchange rate"
	htmlContent := data
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(that.APIKey)
	response, err := client.Send(message)
	if err != nil || !SuccessSentStatusCode(response.StatusCode) {
		fmt.Fprintf(os.Stderr, "error EmailService.Send: %s.(Status code: %d)", err, response.StatusCode)
		return fmt.Errorf("error: %w, status code: %d ", err, response.StatusCode)
	}
	log.Printf("Email sended to %s successfuly.(Status code: %d)", email, response.StatusCode)
	return nil
}

func SuccessSentStatusCode(statusCode int) bool {
	return statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted
}
