package home

import (
	"fmt"
	"net/http"

	"github.com/v3ronez/ufantasyai/handler"
	"github.com/v3ronez/ufantasyai/view/home"
)

func HandlerHomeIndex(w http.ResponseWriter, r *http.Request) error { //is like a controller in laravel
	user := handler.GetAuthenticatedUser(r)
	// acc, err := db.GetAccountUserByID(user.ID)
	// if err != nil {
	// 	return err
	// }
	fmt.Printf("%+v\n", user.Account)
	return home.Index().Render(r.Context(), w) //render view/home func index()
}
