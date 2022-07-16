package midi

import (
	"context"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// subscribeHandler accepts the WebSocket connection and then subscribes
// it to all future messages.
func Play(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"localhost:3000"},
	})
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "something went wrong")

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
	defer cancel()

	for {
		m := midiSignal{}
		err = wsjson.Read(ctx, c, &m)
		if err != nil {
			log.Printf("failed to read with %v: %v", r.RemoteAddr, err)
			return
		}

		log.Printf("\nmidi number: %v, state: %v", m.MIDINum, m.State)
	}
}
