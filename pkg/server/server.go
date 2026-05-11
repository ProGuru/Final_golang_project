package server

import (
	"net/http"
)

const webDir = "./web"

func StartServer() error {
	fs := http.FileServer(http.Dir(webDir))

	http.Handle("/", fs)

	return http.ListenAndServe(":7540", nil)
}
