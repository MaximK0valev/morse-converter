package main

import (
	"log"
	"os"

	"github.com/MaximK0valev/morse-converter/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	srv := server.NewServer(logger)
	err := srv.HTTP.ListenAndServe()
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
