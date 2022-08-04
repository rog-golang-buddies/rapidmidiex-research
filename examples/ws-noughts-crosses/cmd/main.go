package main

import (
	"log"
	"net/http"

	service "ws.rog.noughtscrosses/www"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("listening on port 8888")
	return http.ListenAndServe(":8888", service.New())
}
