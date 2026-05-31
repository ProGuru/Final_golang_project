package api

import (
	"net/http"
	"path/filepath"
)

// Для регистрации API обработчиков

func Init(webDir string) {
	http.Handle("/", http.FileServer(http.Dir(filepath.Clean(webDir))))
	http.HandleFunc("/api/nextdate", nextDayHandler)
}
