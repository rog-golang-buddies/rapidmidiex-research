package main

import (
	"net/http"

	"example.com/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := handlers.Server{}

	r.Route("/time", func(r chi.Router) {
		r.Post("/sync", s.SyncTime)
		r.Get("/stop", s.Stop)
	})
	http.ListenAndServe(":8080", r)
}
