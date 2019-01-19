package controller

import (
	"net/http"
	"encoding/json"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/service/auth"

	"github.com/go-chi/render"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var form auth.LoginForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}


	userI, errors := auth.Login(form)
	if errors == nil {
		byteData, _ := json.Marshal(userI)
		var user *model.User
		json.Unmarshal(byteData, &user)
		render.Render(w, r, user)
	} else {
		render.Render(w, r, ErrNotFound)
	}
}
