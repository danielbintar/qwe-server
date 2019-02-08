package controller

import (
	"net/http"
	"encoding/json"

	"github.com/danielbintar/qwe-server/db"
	"github.com/danielbintar/qwe-server/model"
	characterService "github.com/danielbintar/qwe-server/service/character"

	"github.com/go-chi/render"
)

func GetMyCharacters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUserID := ctx.Value("jwt").(*model.Jwt).UserID

	var characters []*model.Character
	db.DB().Where("user_id = ?", currentUserID).Find(&characters)

	render.RenderList(w, r, SerializeList(characters))
}

func CreateMyCharacter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUserID := ctx.Value("jwt").(*model.Jwt).UserID
	form := characterService.CreateForm{UserID: currentUserID}

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	characterI, errors := characterService.Create(form)
	if errors == nil {
		character := characterI.(*model.Character)
		render.Render(w, r, character.Serialize())
	} else {
		render.Render(w, r, ErrInvalidRequest(errors[0]))
	}
}
