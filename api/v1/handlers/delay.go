package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) DelayHandler(w http.ResponseWriter, r *http.Request) {
	delay, err := strconv.Atoi(chi.URLParam(r, "delay"))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Invalid 'delay' parameter: %s", err.Error()))
		http.Error(w, "Invalid 'delay' parameter", http.StatusBadRequest)

		return
	}

	// Sleep for the specified delay
	time.Sleep(time.Duration(delay) * time.Second)

	// Respond
	fmt.Fprintf(w, "Delayed for %d seconds", delay)
}
