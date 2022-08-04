package http

import (
	"net/http"

	t "github.com/hyphengolang/prelude/template"
)

func (s *Service) routes() {
	s.m.Get("/", s.handleIndex())
	s.m.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	s.m.HandleFunc("/play", chain(s.handleWebSocket(), s.upgradeHTTP()))
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
		go s.newClient(w, r)
	}
}
