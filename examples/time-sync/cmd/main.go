package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rog-golang-buddies/realtime-midi/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := handlers.Server{}

	r.Route("/time", func(r chi.Router) {
		r.Get("/sync", s.SyncTime)
		r.Get("/stop", s.Stop)
	})

	port := "8080"
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}
	fmt.Printf("Starting Clock Sync server on port :%v...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
