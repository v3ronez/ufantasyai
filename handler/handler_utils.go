package handler

import (
	"log/slog"
	"net/http"
)

func MakeHandler(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			slog.Error("Internal server error", "error:", err, "path", r.URL.Path)
		}
	}
}
