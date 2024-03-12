package auth

import (
	"net/http"

	"github.com/v3ronez/ufantasyai/view/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)
}
