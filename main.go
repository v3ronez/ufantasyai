package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/v3ronez/ufantasyai/db"
	"github.com/v3ronez/ufantasyai/handler"
	"github.com/v3ronez/ufantasyai/handler/auth"
	"github.com/v3ronez/ufantasyai/handler/home"
	"github.com/v3ronez/ufantasyai/handler/settings"
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

	router.Get("/login", handler.Make(auth.HandleLoginIndex))
	router.Post("/login", handler.Make(auth.HandleLoginCreate))
	router.Get("/login/provider/google", handler.Make(auth.HandlerLoginWithGoogle))

	router.Post("/logout", handler.Make(auth.HandlerLogout))

	router.Get("/signup", handler.Make(auth.HandleSingUpIndex))
	router.Post("/signup", handler.Make(auth.HandleSingUpCreate))
	router.Get("/auth/redirect-callback", handler.Make(auth.HandlerAuthRedirect)) //redirect after verify email

	router.Group(func(authSetup chi.Router) {
		authSetup.Use(auth.WithAccountSetup)
		authSetup.Get("/account/setup", handler.Make(settings.HandleAccountSetup))
		authSetup.Post("/account/setup", handler.Make(settings.HandlePostAccountSetup))
	})

	//only user logged
	router.Group(func(authRoute chi.Router) {
		// authRoute.Use(auth.WithUserAuth, auth.WithAccountSetup)
		authRoute.Get("/", handler.Make(home.HandlerHomeIndex))
		authRoute.Get("/settings", handler.Make(settings.HandlerSettingsIndex))
		authRoute.Put("/settings", handler.Make(settings.HandlerSettingsUsernameUpdate))

		authRoute.Get("/generate", handler.Make(handler.HandlerGenerateIndex))
		authRoute.Post("/generate", handler.Make(handler.HandlerGenerateImageCreate))
		authRoute.Get("/generate/image/{id}/status", handler.Make(handler.HandlerGenerateImageStatus))
	})

	port := os.Getenv("HTTP_PORT")
	slog.Info("Application running in", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initApp() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	if err := db.Init(); err != nil {
		return err
	}
	return sb.InitSB()
}
