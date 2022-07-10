package http

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"

	t "github.com/topheruk-go/util/template"
)

func (s *Service) routes() {
	s.m.Get("/", s.handleIndex())
	s.m.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	s.m.HandleFunc("/ws", s.handleWebSocket())
}

func (s *Service) handleIndex() http.HandlerFunc {
	render, err := t.Render("pages/index.html")
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, r, nil)
	}
}

func (s *Service) handleWebSocket() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			fmt.Fprintf(w, "upgrade error: %s", err)
			return
		}

		go s.pool.NewConn(c)
	}
}
