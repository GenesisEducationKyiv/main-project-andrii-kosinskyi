package email

import (
	"fmt"
	"net/http"

	"bitcoin_checker_api/internal/model"

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

const htmlTemplate = "Actual exchange rate %s to %s equal %f, this information actual on %s"

func (that *Service) Send(email string, er *model.ExchangeRate) error {
	from := mail.NewEmail(that.FromName, that.FromAddress)
	subject := "Current exchange rate by your subscription"
	to := mail.NewEmail("Dear customer", email)
	plainTextContent := "Current exchange rate"
	htmlContent := fmt.Sprintf(htmlTemplate, er.BaseCurrency, er.QuoteCurrency, er.Price, er.Date)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(that.APIKey)
	response, err := client.Send(message)
	if err != nil || !that.SuccessSentStatusCode(response.StatusCode) {
		return fmt.Errorf("error: %w, status code: %d ", err, response.StatusCode)
	}
	return nil
}

func (that *Service) SuccessSentStatusCode(statusCode int) bool {
	return statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted
}
