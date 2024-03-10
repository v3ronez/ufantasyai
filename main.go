package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initApp(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	// router.Get("/", )
	port := os.Getenv("HTTP_PORT")
	slog.Info("Application running in", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initApp() error {
	// if err := godotenv.Load(); err != nil {
	// retyur
	// }
	return godotenv.Load()
}
