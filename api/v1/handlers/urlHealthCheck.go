package handlers

import (
	"context"
	"core/api/v1/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	url    string
	status string
}

func (h *Handler) URLHealthCheck(w http.ResponseWriter, r *http.Request) {
	var requestBody models.URLHealthCheckRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error decoding the request body: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	defer r.Body.Close()

	h.Logger.Info(fmt.Sprintf("Request body: %+v", requestBody))

	response := make(map[string]string)
	var wg sync.WaitGroup
	resultsChan := make(chan Result, len(requestBody.Links))

	for _, url := range requestBody.Links {
		h.Logger.Info("Ping: " + url)
		wg.Add(1)
		go func(url string, wg *sync.WaitGroup) {
			defer wg.Done()

			// context with a 30-second timeout
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			// HTTP request with the context
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				resultsChan <- Result{url: url, status: "inactive"}
				return
			}

			resp, err := h.HTTPClient.Do(req)
			if err != nil {
				// Check if the error is a timeout
				if ctx.Err() == context.DeadlineExceeded {
					resultsChan <- Result{url: url, status: "timeout"}
				} else {
					resultsChan <- Result{url: url, status: "inactive"}
				}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				resultsChan <- Result{url: url, status: "inactive"}
			} else {
				resultsChan <- Result{url: url, status: "active"}
			}

		}(url, &wg)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(resultsChan)
	}(&wg)

	// Read from the channel as results arrive
	for r := range resultsChan {
		h.Logger.Info(fmt.Sprintf("Ping result: %+v", r))
		response[r.url] = r.status
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error marshalling the response body: %s", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
