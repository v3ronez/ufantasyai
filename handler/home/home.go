package home

import (
	"net/http"

	"github.com/v3ronez/ufantasyai/view/home"
)

func HandlerHomeIndex(w http.ResponseWriter, r *http.Request) error { //is like a controller in laravel
	return home.Index().Render(r.Context(), w) //render view/home
}
