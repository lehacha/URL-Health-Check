package tests

import (
	"bytes"
	"core/api/v1/handlers"
	"core/api/v1/models"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap/zapcore"
)

type MockHTTPClient struct{}

func (c *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Logic to return mocked responses or errors
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString("Response body")),
	}, nil
}

type MockLogger struct{}

func (l *MockLogger) Info(message string, fields ...zapcore.Field)  {}
func (l *MockLogger) Error(message string, fields ...zapcore.Field) {}

func TestURLHealthCheck(t *testing.T) {
	mockClient := &MockHTTPClient{}
	mockLogger := &MockLogger{}
	handler := handlers.NewHandler(mockLogger, mockClient)

	// Create test request body
	requestBody := models.URLHealthCheckRequestBody{
		Links: []string{"http://example.com"},
	}
	requestBodyBytes, _ := json.Marshal(requestBody)
	reqBody := bytes.NewBuffer(requestBodyBytes)

	req := httptest.NewRequest("POST", "http://localhost:8080/pingURLs", reqBody)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the URLHealthCheck function
	handler.URLHealthCheck(rr, req)

	// Check the status code and response body
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	t.Log("TestURLHealthCheck test passed successfully")
}
