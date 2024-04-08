package handler

import (
	"net/http"

	"github.com/v3ronez/ufantasyai/view/credits"
)

func HandleCreditsIndex(w http.ResponseWriter, r *http.Request) error {
	return RenderComponent(w, r, credits.Index())
}
