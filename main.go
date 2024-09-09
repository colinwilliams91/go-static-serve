package main

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// import (
// 	"log"
// 	"mime"
// 	"net/http"

// 	"github.com/colinwilliams91/go-static-serve/internal/handlers"
// 	"github.com/colinwilliams91/go-static-serve/internal/middleware"
// 	"github.com/gorilla/mux"
// )

var compressionExtensions = map[string]string{
    ".gz":  "gzip",
    ".br":  "br",
}

const host = "localhost"
const port = ":8080"

// func main() {
//     router := mux.NewRouter()

//     mime.AddExtensionType(".data", "application/octet-stream")

//     router.HandleFunc("/", handlers.ServeCompressedFiles).Methods("GET")

//     handler := middleware.CacheControlMiddleware(router)

//     log.Printf("Serving files on http://%s%s", host, port)

//     err := http.ListenAndServe(port, handler)

//     if err != nil {
//         log.Fatal(err)
//     }
// }

func main() {
	// Add custom MIME type mapping for .data files (e.g., Unity WebGL data files)
    mime.AddExtensionType(".data", "application/octet-stream")

    // Serve static files
    http.HandleFunc("/", serveCompressedFiles)

    // Start the server
    log.Println("Serving files on http://localhost:8080")
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal(err)
    }
}

// serveCompressedFiles checks if the file has a compressed version (.gz or .br) and serves it with correct headers.
func serveCompressedFiles(w http.ResponseWriter, r *http.Request) {
    requestedPath := r.URL.Path
	log.Printf("Requested Path from URL %s", requestedPath)

    if requestedPath == "/" {
        requestedPath = "/index.html"
    }

    // Try to serve Brotli or Gzip version if available
    for ext, encoding := range compressionExtensions {
		log.Printf("--iterate-- %s -- %s", ext, encoding)

        compressedFilePath := filepath.Join(".", "static", requestedPath+ext)
		log.Printf("COMPRESSED FILE PATH: %s", compressedFilePath)

		if _, err := os.Stat(compressedFilePath); err == nil {
			log.Printf("Request for %s, serving compressed file %s with encoding %s\n", requestedPath, compressedFilePath, encoding)

            w.Header().Set("Content-Encoding", encoding)

			serveFile(w, r, compressedFilePath)

			return
        }
    }

    // Otherwise, serve the file without compression
    toServe := filepath.Join(".", "static", requestedPath)
    log.Printf("~~ serving uncompressed file ~~ %s", toServe)
    serveFile(w, r, toServe)
}

// serveFile serves the requested file and sets the correct Content-Type header.
func serveFile(w http.ResponseWriter, r *http.Request, filePath string) {
    ext := filepath.Ext(filePath)

    // If it's a compressed file, strip the compression extension to get the base MIME type
    if encoding, exists := compressionExtensions[ext]; exists {

        baseFile := strings.TrimSuffix(filePath, ext)

		ext = filepath.Ext(baseFile)

		w.Header().Set("Content-Encoding", encoding)
    }

    // Detect the MIME type based on the file extension
    mimeType := mime.TypeByExtension(ext)

	// if ext == ".html" {
	// 	mimeType = "text/html"
	// }

    if mimeType == "" {
        mimeType = "application/octet-stream"
    }

    w.Header().Set("Content-Type", mimeType)

    // *** DISABLE CACHING *** FOR DEVELOPMENT ONLY? ***
    w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Expires", "0")

    // Serve the file
	log.Printf("Serving file: %s with MIME type: %s\n", filePath, mimeType)
    // err := http.ServeFile(w, r, filePath)
    http.ServeFile(w, r, filePath)
}