package handler

import (
	"log/slog"
	"net/http"

	"github.com/v3ronez/ufantasyai/types"
)

func GetAuthenticatedUser(r *http.Request) types.AuthenticateUser {
	u, ok := r.Context().Value(types.UserContextKey).(types.AuthenticateUser)
	if !ok {
		return types.AuthenticateUser{}
	}
	return u
}

func Make(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			slog.Error("Internal server error", "error:", err, "path", r.URL.Path)
		}
	}
}
