package main

import (
	"fmt"

	"github.com/ProGuru/Final_golang_project/pkg/server"
)

func main() {
	err := server.StartServer()

	if err != nil {
		fmt.Printf("Ошибка запуска сервера: %v", err)
	}
}
