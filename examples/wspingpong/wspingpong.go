package wspingpong

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type LogLevel int

const (
	LogLevelBasic LogLevel = iota
	LogLevelBasicWithHeaders
	LogLevelFull
	LogLevelFullSpew
)

// A service that does some http-hosting, some ws-communication and some logging (with a configurable loglevel)
type WSPingPongServer struct {
	LogLevel LogLevel
}

func (h WSPingPongServer) logBasicRequest(r http.Request) {
	log.Printf("[%s]-request from [%s] using [%s] with [%d] headers\n", r.Method, r.Host, r.Proto, len(r.Header))
}

// Don't call this function! Only call `WSPingPongServer.log`. Otherwise call-depth is not correct.
// TODO: make this function detect if it's called by `WSPingPongServer.log` or something else.
func (h WSPingPongServer) logBasicWithHeadersRequest(r http.Request) {
	s := fmt.Sprintf("[%s]-request from [%s] using [%s] with [%d] headers\n", r.Method, r.Host, r.Proto, len(r.Header))
	log.Output(3, s) // use call-depth 3 to log the line-number of the original calling function
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			fmt.Printf("    %s: %v\n", k, v)
		}
	}
}

// r: not a pointer because we don't want to change the request
func (h WSPingPongServer) log(r http.Request) {
	switch h.LogLevel {
	case LogLevelBasic:
		h.logBasicRequest(r)
	case LogLevelBasicWithHeaders:
		h.logBasicWithHeadersRequest(r)
	case LogLevelFull:
		log.Printf("request: [%+v]\n", r)
	case LogLevelFullSpew:
		log.Println()
		spew.Dump(r)
	}
}

func (h WSPingPongServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	// a manual router ;-)
	switch r.URL.String() {
	case "/favicon.ico":
		return
	case "/redirect":
		w.Header().Add("Content-Type", "") // don't write a html body
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	case "/openandclosewebsocket":
		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		h.log(*r)
		defer c.Close(websocket.StatusInternalError, "the server-sky is falling")

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		var v interface{}
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			log.Println(err)
		}

		log.Printf("received: %v", v)

		c.Close(websocket.StatusNormalClosure, "")
		return
	}

	h.log(*r)

	w.Write([]byte("Hello, I am WSPingPong. This is just plaintext. Nothing more to see here. Maybe try another path?\n"))
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
