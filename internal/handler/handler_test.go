//nolint:goimports // stranger issue on 17 line
package handler_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"bitcoin_checker_api/internal/renderer"

	"bitcoin_checker_api/internal/handler"

	"bitcoin_checker_api/config"
	"bitcoin_checker_api/internal/pkg/email"
	exchangerate "bitcoin_checker_api/internal/pkg/exchange-rate"
	"bitcoin_checker_api/internal/repository"
	"bitcoin_checker_api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHandler_Rate(t *testing.T) {
	repo, err := repository.NewMockRepository(&config.Storage{})
	if err != nil {
		log.Fatal(err)
	}
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_Rate() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_Rate() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	emailServ := email.NewMockService(&config.EmailService{
		APIKey:      "t",
		FromAddress: "t",
		FromName:    "t",
	})

	h := handler.NewHandler(usecase.NewUseCase(repo, excRateCoinpaprika, emailServ), renderer.NewRender())

	r := SetUpRouter()
	r.GET("/", h.Rate)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK || len(w.Body.String()) == 0 {
		t.Errorf("TestHandler_Rate status code: %d  Body: %s", w.Code, w.Body.String())
	}
}

func TestHandler_RateWithError(t *testing.T) {
	repo, err := repository.NewMockRepository(&config.Storage{})
	if err != nil {
		log.Fatal(err)
	}
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{})
	if err != nil {
		t.Errorf("TestHandler_RateWithError() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{})
	if err != nil {
		t.Errorf("TestHandler_RateWithError() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	emailServ := email.NewMockService(&config.EmailService{
		APIKey:      "t",
		FromAddress: "t",
		FromName:    "t",
	})
	h := handler.NewHandler(usecase.NewUseCase(repo, excRateCoinpaprika, emailServ), renderer.NewRender())

	r := SetUpRouter()
	r.GET("/", h.Rate)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest || !strings.Contains(w.Body.String(), handler.ErrInvStatVal) {
		t.Errorf("TestHandler_Rate status code: %d  Body: %s", w.Code, w.Body.String())
	}
}

func TestHandler_Subscription(t *testing.T) {
	repo, err := repository.NewMockRepository(&config.Storage{})
	if err != nil {
		log.Fatal(err)
	}
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_Subscription() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_Subscription() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	emailServ := email.NewMockService(&config.EmailService{
		APIKey:      "t",
		FromAddress: "t",
		FromName:    "t",
	})
	h := handler.NewHandler(usecase.NewUseCase(repo, excRateCoinpaprika, emailServ), renderer.NewRender())

	r := SetUpRouter()
	r.POST("/", h.Subscription)
	data := url.Values{}
	data.Set("email", "taras@schewchs.com")
	postForm := strings.NewReader(data.Encode())
	req, _ := http.NewRequest(http.MethodPost, "/", postForm)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK && w.Body.String() != handler.EmailAdded {
		t.Errorf("TestHandler_Rate status code: %d  Body: %s", w.Code, w.Body.String())
	}
}

func TestHandler_SubscriptionWithError(t *testing.T) {
	repo, err := repository.NewMockRepository(&config.Storage{})
	if err != nil {
		log.Fatal(err)
	}
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SubscriptionWithError() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SubscriptionWithError() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	emailServ := email.NewMockService(&config.EmailService{
		APIKey:      "t",
		FromAddress: "t",
		FromName:    "t",
	})
	h := handler.NewHandler(usecase.NewUseCase(repo, excRateCoinpaprika, emailServ), renderer.NewRender())

	r := SetUpRouter()
	r.POST("/", h.Subscription)
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusConflict && w.Body.String() != handler.ErrInvSubEmail {
		t.Errorf("TestHandler_Rate status code: %d  Body: %s", w.Code, w.Body.String())
	}
}

func TestHandler_SendEmails(t *testing.T) {
	repo, err := repository.NewMockRepository(&config.Storage{})
	if err != nil {
		log.Fatal(err)
	}

	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SendEmails() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SendEmails() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	emailServ := email.NewMockService(&config.EmailService{
		APIKey:      "t",
		FromAddress: "t",
		FromName:    "t",
	})
	h := handler.NewHandler(usecase.NewUseCase(repo, excRateCoinpaprika, emailServ), renderer.NewRender())

	r := SetUpRouter()
	r.POST("/", h.SendEmails)
	r.POST("/subscribe", h.Subscription)

	data := url.Values{}
	data.Set("email", "taras@schewchs.com")
	postForm := strings.NewReader(data.Encode())
	req, _ := http.NewRequest(http.MethodPost, "/subscribe", postForm)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest(http.MethodPost, "/", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK && w.Body.String() != handler.EmailsSent {
		t.Errorf("TestHandler_SendEmails status code: %d  Body: %s", w.Code, w.Body.String())
	}
}

func TestHandler_SendEmailsWithErrorEmptyStorage(t *testing.T) {
	repo, err := repository.NewMockRepository(&config.Storage{})
	if err != nil {
		log.Fatal(err)
	}
	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SendEmailsWithErrorEmptyStorage() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SendEmailsWithErrorEmptyStorage() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	emailServ := email.NewMockService(&config.EmailService{
		APIKey:      "t",
		FromAddress: "t",
		FromName:    "t",
	})
	h := handler.NewHandler(usecase.NewUseCase(repo, excRateCoinpaprika, emailServ), renderer.NewRender())

	r := SetUpRouter()
	r.POST("/", h.SendEmails)

	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusConflict && strings.Contains(w.Body.String(), handler.ErrEmailsNotSent) {
		t.Errorf("TestHandler_SendEmailsWithErrorEmptyStorage status code: %d  Body: %s", w.Code, w.Body.String())
	}
}

func TestHandler_SendEmailsWithErrorExchangeRateService(t *testing.T) {
	repo, err := repository.NewMockRepository(&config.Storage{})
	if err != nil {
		log.Fatal(err)
	}

	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SendEmailsWithErrorExchangeRateService() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SendEmailsWithErrorExchangeRateService() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	emailServ := email.NewMockService(&config.EmailService{
		APIKey:      "t",
		FromAddress: "t",
		FromName:    "t",
	})
	h := handler.NewHandler(usecase.NewUseCase(repo, excRateCoinpaprika, emailServ), renderer.NewRender())

	r := SetUpRouter()
	r.POST("/", h.SendEmails)
	r.POST("/subscribe", h.Subscription)

	data := url.Values{}
	data.Set("email", "taras@schewchs.com")
	postForm := strings.NewReader(data.Encode())
	req, _ := http.NewRequest(http.MethodPost, "/subscribe", postForm)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest(http.MethodPost, "/", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusConflict && strings.Contains(w.Body.String(), handler.ErrEmailsNotSent) {
		t.Errorf("TestHandler_SendEmailsWithErrorExchangeRateService status code: %d  Body: %s", w.Code, w.Body.String())
	}
}

func TestHandler_SendEmailsWithErrorInEmailService(t *testing.T) {
	repo, err := repository.NewMockRepository(&config.Storage{})
	if err != nil {
		log.Fatal(err)
	}

	excRateCoinpaprika, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SendEmailsWithErrorInEmailService() err = %v", err)
	}
	excRateBinance, err := exchangerate.NewMockExchangeRate(&config.DefaultExchangeRate{
		URLMask: "https://api.coinpaprika.com/v1/price-converter?base_currency_id=%s&quote_currency_id=%s&amount=1",
		InRate:  "btc-bitcoin",
		OutRate: "uah-ukrainian-hryvnia",
	})
	if err != nil {
		t.Errorf("TestHandler_SendEmailsWithErrorInEmailService() err = %v", err)
	}
	excRateCoinpaprika.SetNext(excRateBinance)
	emailServ := email.NewMockService(&config.EmailService{
		APIKey:      "",
		FromAddress: "",
		FromName:    "",
	})
	h := handler.NewHandler(usecase.NewUseCase(repo, excRateCoinpaprika, emailServ), renderer.NewRender())

	r := SetUpRouter()
	r.POST("/", h.SendEmails)
	r.POST("/subscribe", h.Subscription)

	data := url.Values{}
	data.Set("email", "taras@schewchs.com")
	postForm := strings.NewReader(data.Encode())
	req, _ := http.NewRequest(http.MethodPost, "/subscribe", postForm)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest(http.MethodPost, "/", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusConflict && strings.Contains(w.Body.String(), handler.ErrEmailsNotSent) {
		t.Errorf("TestHandler_SendEmailsWithErrorInEmailService status code: %d  Body: %s", w.Code, w.Body.String())
	}
}
