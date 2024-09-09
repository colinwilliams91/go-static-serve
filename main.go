package main

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var compressionExtensions = map[string]string{
    ".gz":  "gzip",
    ".br":  "br",
}


func main() {
	// Add custom MIME type mapping for .data files (e.g., Unity WebGL data files)
    mime.AddExtensionType(".data", "application/octet-stream")

    // Serve static files
    http.HandleFunc("/", serveCompressedFiles)

    // Start the server
    log.Println("Serving files on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

// serveCompressedFiles checks if the file has a compressed version (.gz or .br) and serves it with correct headers.
func serveCompressedFiles(w http.ResponseWriter, r *http.Request) {
    // Get the requested file path
    requestedPath := r.URL.Path
	log.Printf("Requested Path from URL %s", requestedPath)

    // Try to serve Brotli or Gzip version if available
    for ext, encoding := range compressionExtensions {
        compressedFilePath := filepath.Join(".", "static", requestedPath+ext)
		log.Printf("COMPRESSED FILE PATH: %s", compressedFilePath)
        if _, err := os.Stat(compressedFilePath); err == nil {
			log.Printf("Request for %s, serving compressed file %s with encoding %s\n", requestedPath, compressedFilePath, encoding)
			// If compressed version exists, serve it with correct encoding
            w.Header().Set("Content-Encoding", encoding)
            serveFile(w, r, compressedFilePath)
            return
        }
    }

    // Otherwise, serve the file without compression
    serveFile(w, r, filepath.Join(".", "static", requestedPath))
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
    if mimeType == "" {
        mimeType = "application/octet-stream"
    }
    w.Header().Set("Content-Type", mimeType)

    // Serve the file
    http.ServeFile(w, r, filePath)
}