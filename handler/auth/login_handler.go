package auth

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/nedpals/supabase-go"
	"github.com/v3ronez/ufantasyai/handler"
	"github.com/v3ronez/ufantasyai/pkg/sb"
	"github.com/v3ronez/ufantasyai/view/auth"
)

const (
	SessionUserKey         = "user"
	SessionAccessTokenName = "accessToken"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)
}

func HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	res, err := sb.Client.Auth.SignIn(r.Context(), credentials)
	if err != nil {
		slog.Error("error", err)
		return handler.RenderComponent(w, r,
			auth.LoginForm(credentials,
				auth.LoginErrors{InvalidCredentials: "The credentials you have entered are invalid"}))
	}

	if err := setSession(w, r, res.AccessToken); err != nil {
		return err
	}
	// setAuthCokkie(w, &http.Cookie{Name: "access_token", Value: res.AccessToken})
	// http.Redirect(w, r, "/", http.StatusSeeOther)
	handler.HtmxRedirect(w, r, "/")
	return nil
}

func HandlerLoginWithGoogle(w http.ResponseWriter, r *http.Request) error {
	resp, err := sb.Client.Auth.SignInWithProvider(supabase.ProviderSignInOptions{
		Provider:   "google",
		RedirectTo: "http://localhost:3000/auth/redirect-callback",
	})
	if err != nil {
		return err
	}
	http.Redirect(w, r, resp.URL, http.StatusSeeOther)
	return nil
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
	//validade fields

	// if !util.IsValidEmail(credentials.Email) {
	// 	return handler.RenderComponent(w, r, auth.LoginForm(credentials, auth.LoginErrors{Email: "Email invalid"}))
	// }
	// if p, ok := util.ValidatePassword(credentials.Password); !ok {
	// 	return handler.RenderComponent(w, r, auth.LoginForm(credentials, auth.LoginErrors{Password: p}))
	// }

	user, err := sb.Client.Auth.SignUp(r.Context(), credentials)
	if err != nil {
		slog.Error("Error to create a new user: ", err)
		return nil
	}
	return handler.RenderComponent(w, r, auth.SignUpSuccess(user.Email))
}

// redirect confirm email link
func HandlerAuthRedirect(w http.ResponseWriter, r *http.Request) error {
	accessToken := r.URL.Query().Get("access_token")
	if accessToken == "" {
		return handler.RenderComponent(w, r, auth.RedictCallBackScript())
	}
	if err := setSession(w, r, accessToken); err != nil {
		return err
	}
	// setAuthCokkie(w, &http.Cookie{Name: "access_token", Value: accessToken})
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

// func setAuthCokkie(w http.ResponseWriter, cokkie *http.Cookie) {
// 	cokkie.HttpOnly = true
// 	cokkie.Secure = true
// 	if cokkie.Expires.IsZero() {
// 		cokkie.Expires = time.Now().Add(time.Hour)
// 	}
// 	if cokkie.Path == "" {
// 		cokkie.Path = "/"
// 	}
// 	http.SetCookie(w, cokkie)
// }

func setSession(w http.ResponseWriter, r *http.Request, token string) error {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, _ := store.Get(r, SessionUserKey)
	session.Values[SessionAccessTokenName] = token
	session.Options.HttpOnly = true
	session.Options.Secure = true
	return session.Save(r, w)
}
func deleteSession(w http.ResponseWriter, r *http.Request) error {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, _ := store.Get(r, SessionUserKey)
	session.Values[SessionAccessTokenName] = ""
	session.Options.MaxAge = -1
	return session.Save(r, w)
}

func HandlerLogout(w http.ResponseWriter, r *http.Request) error {
	// setAuthCokkie(w, &http.Cookie{
	// 	Name:   "access_token",
	// 	Value:  "",
	// 	MaxAge: -1,
	// })
	if err := deleteSession(w, r); err != nil {
		return err
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}
