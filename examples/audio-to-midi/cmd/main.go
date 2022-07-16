package main

import (
	"fmt"
	"log"
	"net/http"

	atm "github.com/rog-golang-buddies/realtime-midi/examples/audio-to-midi"
)

func main() {
	server := atm.Server{
		Port:   ":8080",
		Router: http.ServeMux{},
	}

	fmt.Printf("starting server on port: %v", server.Port)
	err := server.ServeHTTP()
	if err != nil {
		log.Fatalf("failed to start server: %v", err.Error())
	}
}
