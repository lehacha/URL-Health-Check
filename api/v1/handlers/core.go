package handlers

import (
	"net/http"

	"go.uber.org/zap/zapcore"
)

type Handler struct {
	Logger     LoggerInterface
	HTTPClient HTTPClientInterface
}

// HTTPClientInterface allows for mocking http.Client
type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

// LoggerInterface allows for mocking a logger
type LoggerInterface interface {
	Info(string, ...zapcore.Field)
	Error(string, ...zapcore.Field)
}

func NewHandler(logger LoggerInterface, httpClient HTTPClientInterface) *Handler {
	return &Handler{
		Logger:     logger,
		HTTPClient: httpClient,
	}
}
