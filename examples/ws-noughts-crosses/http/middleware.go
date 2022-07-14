package http

import (
	"context"
	"fmt"
	"net/http"

	// h "ws-chat.example/pkg/http"

	h "ws.rog.noughtscrosses/pkg/http"

	"github.com/gorilla/websocket"
)

var upgradeKey = &contextKey{"upgrade-http"}

type MiddleWare func(f http.HandlerFunc) http.HandlerFunc

func chain(f http.HandlerFunc, mw ...MiddleWare) http.HandlerFunc {
	for _, m := range mw {
		f = m(f)
	}
	return f
}

func (s *Service) upgradeHTTP() MiddleWare {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, w.Header())
			if err != nil {
				h.Respond(w, r, err, http.StatusForbidden)
				return
			}

			fmt.Println("handle new request")
			f(w, r.WithContext(context.WithValue(r.Context(), upgradeKey, conn)))
		}
	}
}
