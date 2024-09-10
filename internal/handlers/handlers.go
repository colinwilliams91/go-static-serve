package handlers

import (
	"log"
	"mime"
	"net/http"

	"path/filepath"
	"strings"
)

var compressionExtensions = map[string]string{
    ".gz":  "gzip",
    ".br":  "br",
}

// TODO: add endpoint (key) target-file (values) as needed...
var endpointHTML = map[string]string {
    "/": "/index.html",
}

/*
-- ==========================================
-- -- server functions					-- --
-- ==========================================
*/

// ServeCompressedFiles checks if the file has a compressed version (.gz or .br) and serves it with correct headers.
func ServeCompressedFiles(w http.ResponseWriter, r *http.Request) {
    requestedPath := r.URL.Path
	log.Printf("Requested Path from URL %s", requestedPath)

    targetHTML, ok := endpointHTML[requestedPath]
    if ok {
        requestedPath = targetHTML
    }

    // TODO: Try to serve Brotli or Gzip version if available if compressed assets become large in PROD (see content-encoding below...)

    // Otherwise, serve the file without compression
    requestedFile := filepath.Join(".", "static", requestedPath)

    serveFile(w, r, requestedFile)
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

    // TODO: add fallback for empty mimeType from ext as needed

    w.Header().Set("Content-Type", mimeType)

	// log.Printf("Serving file: %s with MIME type: %s\n", filePath, mimeType)

    http.ServeFile(w, r, filePath)
}

/* -- FALLBACK MIME TYPE --

    if mimeType == "" {
        mimeType = "application/octet-stream"
    }

*/

/* -- MANUAL COMPRESSION CONTENT-ENCODING --

    for ext, encoding := range compressionExtensions {

        baseFileName := strings.TrimSuffix(requestedPath, filepath.Ext(requestedPath))

        compressedFilePath := filepath.Join(".", "static", baseFileName+ext)

        log.Printf(compressedFilePath)

		if _, err := os.Stat(compressedFilePath); err == nil {
			log.Printf("Request for %s, serving compressed file %s with encoding %s\n", requestedPath, compressedFilePath, encoding)

            w.Header().Set("Content-Encoding", encoding)

			serveFile(w, r, compressedFilePath)

			return
        }
    }
*/
