package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type syncReq struct {
	BPM int64 `json:"bpm"`
}

type Server struct {
	WSConn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	done   = make(chan bool)
	ticker *time.Ticker
)

// start syncing time
func (s *Server) SyncTime(w http.ResponseWriter, r *http.Request) {
	var err error

	// create websocket connection
	s.WSConn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// create Server
	sr := syncReq{}

	// read request data
	err = s.WSConn.ReadJSON(&sr)
	if err != nil {
		log.Println(err)
		return
	}

	bd := time.Duration(time.Minute.Milliseconds() / sr.BPM * int64(time.Millisecond)) // beat duration

	// start MIDI ticker, works like a metronome
	ticker = time.NewTicker(bd)

	// check for ticks in background
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				s.WSConn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			}
		}
	}()
}

// stop ticking and close websocket connection
func (s *Server) Stop(w http.ResponseWriter, r *http.Request) {
	ticker.Stop()
	done <- true
	err := s.WSConn.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stopped"))
}
