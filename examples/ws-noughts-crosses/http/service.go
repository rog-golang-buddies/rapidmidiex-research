package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/rog-golang-buddies/realtime-midi/examples/ws-noughts-crosses/http/ws"
)

type Service struct {
	m    *chi.Mux
	pool *ws.Pool
}

func New() *Service {
	s := &Service{
		m:    chi.NewMux(),
		pool: ws.NewPool(),
	}
	s.routes()
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.m.ServeHTTP(w, r)
}
