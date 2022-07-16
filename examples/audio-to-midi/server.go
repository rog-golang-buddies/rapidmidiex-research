package atm

import (
	"log"
	"net/http"
	"time"

	"github.com/rog-golang-buddies/realtime-midi/examples/audio-to-midi/midi"
	"github.com/rs/cors"
)

type Server struct {
	Port   string
	Router http.ServeMux
}

func (s *Server) initRoutes() {
	s.Router.HandleFunc("/midi/play", midi.Play)
}

func (s *Server) ServeHTTP() error {
	s.initRoutes()

	c := cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}
	handler := cors.New(c).Handler(&s.Router)

	server := http.Server{
		Addr:         s.Port,
		Handler:      handler,
		ErrorLog:     log.Default(),     // set the logger for the server
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
