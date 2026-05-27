package server

import (
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logFile *log.Logger, webDir string) *Server {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepath.Clean(webDir))))

	return &Server{
		Logger: logFile,
		Server: &http.Server{
			Addr:         ":7540",
			Handler:      mux,
			ErrorLog:     logFile,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}
