package main

import (
	"log"
	"net/http"

	// websocket "ws.rog.noughtscrosses"

	service "ws.rog.noughtscrosses/http"
)

func main() {
	log.Println("listening on port 8081")
	http.ListenAndServe(":8081", service.New())
}
