package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"ws.rog.numberguesser/net"
)

type Service struct {
	m    *chi.Mux
	pool *net.Pool
}

func New() *Service {
	s := &Service{
		m:    chi.NewMux(),
		pool: net.NewPool(),
	}
	s.routes()
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.m.ServeHTTP(w, r)
}
