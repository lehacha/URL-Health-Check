package main

import (
	"core/internal/logger"
	"fmt"
	"net/http"

	"core/api/v1/handlers"
	v1Middleware "core/api/v1/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Default().Handler)

	logger := logger.NewLogger()
	logger.Info("Logger created")

	r.Use(v1Middleware.ZapLogger(logger))

	// API
	handler := handlers.NewHandler(logger, http.DefaultClient)
	r.Post("/pingURLs", handler.URLHealthCheck)
	r.Get("/delay/{delay}", handler.DelayHandler)

	port := "8080"
	logger.Info(fmt.Sprintf("Starting the service on the port %s", port))
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		logger.Panic("Can not start service: " + err.Error())
	}

}
