package handler

import (
	"net/http"

	"github.com/v3ronez/ufantasyai/view/generate"
)

func HandlerGenerateIndex(w http.ResponseWriter, r *http.Request) error {
	return RenderComponent(w, r, generate.Index())
}
