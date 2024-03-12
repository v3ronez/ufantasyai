package auth

import (
	"log/slog"
	"net/http"

	"github.com/nedpals/supabase-go"
	"github.com/v3ronez/ufantasyai/handler"
	"github.com/v3ronez/ufantasyai/pkg/sb"
	"github.com/v3ronez/ufantasyai/view/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)
}

func HandleLoginUser(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	// if !util.IsValidEmail(credentials.Email) {
	// 	return handler.RenderComponent(w, r, auth.LoginForm(credentials, auth.LoginErrors{Email: "Email invalid"}))
	// }
	// if p, ok := util.ValidatePassword(credentials.Password); !ok {
	// 	return handler.RenderComponent(w, r, auth.LoginForm(credentials, auth.LoginErrors{Password: p}))
	// }
	res, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	if err != nil {
		slog.Error("error", err)
		return handler.RenderComponent(w, r,
			auth.LoginForm(credentials,
				auth.LoginErrors{InvalidCredentials: "The credentials you have entered are invalid"}))
	}
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    res.AccessToken,
		Secure:   true,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusOK)
	return nil
	// return handler.RenderComponent(w, r, auth.LoginForm(credentials, auth.LoginErrors{}))
}

// sing up
func HandleSingUpIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.SignUp().Render(r.Context(), w)
}

func HandleSingUpCreate(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	_, err := sb.Client.Auth.SignUp(r.Context(), credentials)
	if err != nil {
		slog.Error("Error to create a new user: ", err)
	}

	return nil
}
