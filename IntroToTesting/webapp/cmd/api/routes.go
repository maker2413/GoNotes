package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// Register middleware
	mux.Use(middleware.Recoverer)
	// mux.Use(app.enableCORS)

	// Authentication routes - auth handler, refresh
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh", app.refresh)

	// Test handler
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		var payload = struct {
			Message string `json:"message"`
		}{
			Message: "hello, world",
		}

		_ = app.writeJSON(w, http.StatusOK, payload)
	})

	// Protected routes
	mux.Route("/users", func(mux chi.Router) {
		// Use auth middleware

		mux.Get("/", app.allUsers)
		mux.Get("/{userID}", app.getUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/{userID}", app.insertUser)
		mux.Patch("/{userID}", app.updateUser)
	})

	return mux
}
