package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
)

func main() {
	server := Server{
		Port: ":8080",
	}

	server.ServeHTTP()
}

type Server struct {
	Port   string
	Router chi.Mux
}

func (s *Server) initRoutes() {
	s.Router.Get("/handler", sayHello)
	s.Router.Get("/handlerf", sayHello2().ServeHTTP)
}

func (s *Server) ServeHTTP() {
	s.initRoutes()

	// don't use "*" as AllowedOrigins, new origins should be added explicitly
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	c := cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}
	handler := cors.New(c).Handler(r)

	serverCtx, serverStop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer serverStop()

	server := http.Server{
		Addr:         s.Port,
		Handler:      handler,
		ErrorLog:     log.Default(),     // set the logger for the server
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		BaseContext: func(_ net.Listener) context.Context {
			return serverCtx
		},
	}

	g, gCtx := errgroup.WithContext(serverCtx)
	g.Go(func() error {
		// Run the server
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return server.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \n", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from Handler"))
}

func sayHello2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world from HandlerFunc"))
	}
}
