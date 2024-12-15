package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	// Create router
	mux := chi.NewRouter()

	// Set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// Define application routes
	mux.Get("/", app.HomePage)

	return mux
}
