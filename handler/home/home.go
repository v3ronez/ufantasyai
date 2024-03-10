package handler

import (
	"net/http"

	"github.com/v3ronez/ufantasyai/view/home"
)

func HandlerHomeIndex(w http.ResponseWriter, r *http.Request) { //is like a controller in laravel
	home.Index().Render(r.Context(), w)
}
