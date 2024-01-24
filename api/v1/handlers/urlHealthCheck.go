package handlers

import (
	"core/api/v1/models"
	"encoding/json"
	"fmt"
	"net/http"
)

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
	response["link1"] = "active"
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
