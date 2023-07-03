package email

import (
	"fmt"
	"net/http"

	"bitcoin_checker_api/config"
)

type MockService struct {
	APIKey      string
	FromAddress string
	FromName    string
}

func NewMockService(c *config.EmailService) *MockService {
	return &MockService{
		APIKey:      c.APIKey,
		FromAddress: c.FromAddress,
		FromName:    c.FromName,
	}
}

func (that *MockService) Send(_, _ string) error {
	if that.APIKey == "" {
		return fmt.Errorf("error: not send, status code: %d ", http.StatusBadRequest)
	}
	return nil
}

func (that *MockService) SuccessSentStatusCode(statusCode int) bool {
	return statusCode == http.StatusOK || statusCode == http.StatusCreated || statusCode == http.StatusAccepted
}
