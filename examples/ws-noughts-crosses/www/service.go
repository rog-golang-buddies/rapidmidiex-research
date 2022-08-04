package www

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"

	"ws.rog.noughtscrosses/www/ws"
)

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "net/http context value " + k.name }

type Service struct {
	m *chi.Mux
	p *ws.Pool
}

func New() *Service {
	s := &Service{
		m: chi.NewMux(),
		p: ws.NewPool(),
	}
	s.routes()
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.m.ServeHTTP(w, r)
}

func (s *Service) newClient(w http.ResponseWriter, r *http.Request) error {
	log.Println("new client")
	return ws.NewClient(s.p, r.Context().Value(upgradeKey).(*websocket.Conn))
}
