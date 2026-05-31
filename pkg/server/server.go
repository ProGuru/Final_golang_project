package server

import (
	"log"
	"net/http"
	"time"

	"github.com/ProGuru/Final_golang_project/pkg/api"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func Run(logFile *log.Logger, webDir string) *Server {
	api.Init(webDir)

	return &Server{
		Logger: logFile,
		Server: &http.Server{
			Addr:         ":7540",
			ErrorLog:     logFile,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}
