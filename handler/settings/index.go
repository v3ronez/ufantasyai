package settings

import (
	"net/http"

	"github.com/v3ronez/ufantasyai/db"
	"github.com/v3ronez/ufantasyai/handler"
	"github.com/v3ronez/ufantasyai/types"
	"github.com/v3ronez/ufantasyai/view/auth"
	"github.com/v3ronez/ufantasyai/view/settings"
)

func HandlerSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	u := handler.GetAuthenticatedUser(r)

	handler.RenderComponent(w, r, settings.Index(u))
	return nil
}

func HandleAccountSetup(w http.ResponseWriter, r *http.Request) error {
	return handler.RenderComponent(w, r, auth.AccountSetup())
}

func HandlePostAccountSetup(w http.ResponseWriter, r *http.Request) error {
	params := auth.AccountSetupFormParams{
		Username: r.FormValue("username"),
	}
	var errs = auth.AccountSetupErrors{}
	if len(params.Username) < 5 {
		errs.Username = "invalid username"
		return handler.RenderComponent(w, r, auth.AccountSetupForm(params, errs))
	}
	user := handler.GetAuthenticatedUser(r)
	acc := types.Account{
		UserId:   user.ID,
		UserName: params.Username,
	}
	if err := db.CreateNewAccount(&acc); err != nil {
		return err
	}

	return handler.HtmxRedirect(w, r, "/")
}
