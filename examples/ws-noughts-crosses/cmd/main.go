package main

import (
	"log"
	"net/http"

	ptth "github.com/rog-golang-buddies/realtime-midi/examples/ws-noughts-crosses/http"
)

func main() {
	log.Println("listening on port 8080")
	http.ListenAndServe(":8080", ptth.New())
}
