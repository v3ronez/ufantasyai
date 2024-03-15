package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/v3ronez/ufantasyai/handler"
	"github.com/v3ronez/ufantasyai/pkg/sb"
	"github.com/v3ronez/ufantasyai/types"
)

// check user logged and create a context
func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		// u, err := userIsLogged(r)
		// if err != nil {
		// 	next.ServeHTTP(w, r)
		// 	return
		// }
		store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
		session, err := store.Get(r, SessionUserKey)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		acccessToken := session.Values[SessionAccessTokenName]
		if acccessToken == nil {
			next.ServeHTTP(w, r)
			return
		}

		resp, err := sb.Client.Auth.User(r.Context(), acccessToken.(string))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		user := types.AuthenticateUser{
			Email:    resp.Email,
			LoggedIn: true,
		}
		ctx := context.WithValue(r.Context(), types.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx)) //forward the request with the user in context
	}
	return http.HandlerFunc(fn)
}

// func userIsLogged(r *http.Request) (types.AuthenticateUser, error) {
// 	c, err := r.Cookie("access_token")
// 	if err != nil {
// 		return types.AuthenticateUser{}, err
// 	}
// 	user, err := sb.Client.Auth.User(r.Context(), c.Value)
// 	if err != nil {
// 		return types.AuthenticateUser{}, err
// 	}
// 	return types.AuthenticateUser{
// 		Email:    user.Email,
// 		LoggedIn: true,
// 	}, nil
// }

func WithUserAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		u := handler.GetAuthenticatedUser(r)
		if !u.LoggedIn {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r) //ServeHTTP call a function to forward the w,r
	}
	return http.HandlerFunc(fn)
}
