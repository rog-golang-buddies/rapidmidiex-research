package main

import (
	"log"
	"net/http"

	ptth "ws.rog.numberguesser/http"
)

func main() {
	log.Println("listening on port 8080")
	http.ListenAndServe(":8080", ptth.New())
}
