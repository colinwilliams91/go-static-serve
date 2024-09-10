package main

import (
	"log"
	"mime"
	"net/http"

	"github.com/colinwilliams91/go-static-serve/internal/handlers"
	"github.com/colinwilliams91/go-static-serve/internal/middleware"
)

const protocol = "http"
const host = "localhost"
const port = ":8080"

func main() {
    mime.AddExtensionType(".data", "application/octet-stream")

    http.HandleFunc("/", handlers.ServeCompressedFiles)

    // TODO: RM in PROD...
    handler := middleware.CacheControlMiddleware(http.DefaultServeMux)

    log.Printf("Serving files on %s://%s%s", protocol, host, port)

    err := http.ListenAndServe(port, handler)

    if err != nil {
        log.Fatal(err)
    }
}
