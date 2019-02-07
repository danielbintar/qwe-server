package controller

import (
	"net/http"

	"github.com/danielbintar/qwe-server/db"
	"github.com/danielbintar/qwe-server/model"

	"github.com/go-chi/render"
)

func GetMyCharacters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUserID := ctx.Value("jwt").(*model.Jwt).UserID

	var characters []*model.Character
	db.DB().Where("user_id = ?", currentUserID).Find(&characters)

	render.RenderList(w, r, SerializeList(characters))
}
