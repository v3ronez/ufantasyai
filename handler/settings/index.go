package settings

import (
	"net/http"

	"github.com/v3ronez/ufantasyai/handler"
	"github.com/v3ronez/ufantasyai/view/settings"
)

func HandlerSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	u := handler.GetAuthenticatedUser(r)

	handler.RenderComponent(w, r, settings.Index(u))
	return nil
}
