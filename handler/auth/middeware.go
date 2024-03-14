package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/v3ronez/ufantasyai/types"
)

func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		user := types.AuthenticateUser{
			// Email:    "test@gmail.com",
			// LoggedIn: true,
		}
		ctx := context.WithValue(r.Context(), types.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx)) //forward the request
	}
	return http.HandlerFunc(fn)
}
