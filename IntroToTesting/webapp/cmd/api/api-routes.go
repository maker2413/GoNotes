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
	mux.Use(app.enableCORS)

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))

	mux.Route(("/web"), func(mux chi.Router) {
		mux.Post("/auth", app.authenticate)
		mux.Get("/refresh-token", app.refreshUsingCookie)
		mux.Get("/logout", app.deleteRefreshCookie)
	})

	// Authentication routes - auth handler, refresh
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh-token", app.refresh)

	// Protected routes
	mux.Route("/users", func(mux chi.Router) {
		// Use auth middleware
		mux.Use(app.authRequired)

		mux.Get("/", app.allUsers)
		mux.Get("/{userID}", app.getUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/{userID}", app.insertUser)
		mux.Patch("/{userID}", app.updateUser)
	})

	return mux
}
