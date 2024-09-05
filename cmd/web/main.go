package main

import (
	"log"
	"net/http"
)

func main() {
	var mux http.Handler = routes()

	log.Println("Starting web server on port 8080...")

	_ = http.ListenAndServe(":8080", mux)
}