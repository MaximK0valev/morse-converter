package server

import (
	"log"
	"net/http"
	"time"

	"github.com/MaximK0valev/morse-converter/internal/handlers"
)

// Server bundles application dependencies and the underlying HTTP server.
type Server struct {
	// Logger is used by the HTTP server and application components for logging.
	Logger *log.Logger

	// HTTP is the configured net/http server instance.
	HTTP *http.Server
}

// NewServer constructs a Server with all routes registered and reasonable timeouts set.
// The returned Server is ready to be started via srv.HTTP.ListenAndServe.
func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	httpserver := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Logger: logger,
		HTTP:   httpserver,
	}
}
