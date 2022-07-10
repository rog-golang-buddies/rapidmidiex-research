package wspingpong

import (
	"log"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type LogLevel int

const (
	LogLevelBasic LogLevel = iota
	LogLevelFull
	LogLevelFullSpew
)

type WSPingPongServer struct {
	LogLevel LogLevel
}

// r: not a pointer because we don't want to change the request
func (h WSPingPongServer) log(r http.Request) {
	switch h.LogLevel {
	case LogLevelBasic:
		log.Printf("request from [%s]\n", r.Host)
	case LogLevelFull:
		log.Printf("request: [%+v]\n", r)
	case LogLevelFullSpew:
		log.Println()
		spew.Dump(r)
	}
}

func (h WSPingPongServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/favicon.ico" {
		return
	}
	h.log(*r)
	w.Write([]byte("Hello, I am WSPingPong\n"))
}

func StartServer(port string, loglevel LogLevel) {
	// Define our own custom server with
	// - custom LogLevel
	// - ...
	s := http.Server{
		Addr:         port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  100 * time.Second,
		Handler: WSPingPongServer{
			LogLevel: loglevel,
		},
	}

	log.Printf("Starting server: %+v", s)
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatalln(err)
		}
		log.Println(err)
	}
}
