package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/v3ronez/ufantasyai/handler"
	"github.com/v3ronez/ufantasyai/handler/auth"
	"github.com/v3ronez/ufantasyai/handler/home"
	"github.com/v3ronez/ufantasyai/pkg/sb"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initApp(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Use(auth.WithUser)
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS)))) //render static files
	// router.Handle("/*", http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))) //render static files
	router.Get("/", handler.Make(home.HandlerHomeIndex))
	router.Get("/login", handler.Make(auth.HandleLoginIndex))
	router.Get("/signup", handler.Make(auth.HandleSingUpIndex))
	router.Post("/signup", handler.Make(auth.HandleSingUpCreate))
	router.Post("/login", handler.Make(auth.HandleLoginUser))

	port := os.Getenv("HTTP_PORT")
	slog.Info("Application running in", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initApp() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return sb.InitSB()
}
