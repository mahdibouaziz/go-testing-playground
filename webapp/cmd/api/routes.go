package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	// mux.Use(app.enableCORS)

	// authentication routes - auth & refresh
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh-token", app.refresh)

	// test handler
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		var payload = struct {
			Message string `json:"message"`
		}{
			Message: "Hello, world",
		}

		_ = app.writeJSON(w, http.StatusOK, payload)
	})

	// protected routes
	mux.Route("/users", func(mux chi.Router) {
		// use auth middleware

		mux.Get("/", app.allUsers)
		mux.Get("/{userId}", app.getUser)
		mux.Delete("/{userId}", app.deleteUser)
		mux.Put("/", app.insertUser)
		mux.Patch("/", app.updateUser)
	})

	return mux
}
