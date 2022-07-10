package main

import (
	"log"
	"net/http"
	"time"
)

type LogLevel int

const (
	Basic LogLevel = iota
	Full
)

type WSPingPongHandler struct {
	LogLevel LogLevel
}

// r: not a pointer because we want to change the request
func (h WSPingPongHandler) log(r http.Request) {
	switch h.LogLevel {
	case Basic:
		log.Printf("request from [%s]\n", r.Host)
	case Full:
		log.Printf("request: [%+v]\n", r)
	}
}

func (h WSPingPongHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/favicon.ico" {
		return
	}
	h.log(*r)
	w.Write([]byte("Hello, I am WSPingPong\n"))
}

func main() {
	// Make log print a datetime and a filename:linenumber
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Define our own custom server with
	// - custom LogLevel
	// - ...
	s := http.Server{
		Addr:         ":9876",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  100 * time.Second,
		Handler: WSPingPongHandler{
			LogLevel: Full,
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
