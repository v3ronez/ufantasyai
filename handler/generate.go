package handler

import (
	"net/http"

	"github.com/v3ronez/ufantasyai/types"
	"github.com/v3ronez/ufantasyai/view/generate"
)

func HandlerGenerateIndex(w http.ResponseWriter, r *http.Request) error {
	images := []types.Image{}
	viewData := generate.ViewData{Images: images}
	return RenderComponent(w, r, generate.Index(viewData))
}
