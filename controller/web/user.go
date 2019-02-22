package web

import (
	"net/http"
	"encoding/json"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/service/user"

	"github.com/go-chi/render"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var form user.CreateForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	userI, errors := user.Create(form)
	if errors == nil {
		byteData, _ := json.Marshal(userI)
		var user *model.User
		json.Unmarshal(byteData, &user)

		token := createToken(user)
		render.Render(w, r, token)
	} else {
		render.Render(w, r, ErrInvalidRequest(errors[0]))
	}
}
