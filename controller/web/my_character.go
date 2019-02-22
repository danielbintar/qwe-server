package web

import (
	"context"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/danielbintar/qwe-server/db"
	"github.com/danielbintar/qwe-server/model"
	characterService "github.com/danielbintar/qwe-server/service/character"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func MyCharacter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		u64, err := strconv.ParseUint(chi.URLParam(r, "characterID"), 10, 32)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		characterID := uint(u64)

		character := &model.Character{ID: characterID}
		db.DB().Where(&character).First(&character)

		if character.Name == "" {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		if character.UserID != ctx.Value("jwt").(*model.Jwt).UserID {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		ctx = context.WithValue(ctx, "character", character)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetMyCharacters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUserID := ctx.Value("jwt").(*model.Jwt).UserID

	var characters []*model.Character
	db.DB().Where("user_id = ?", currentUserID).Find(&characters)

	render.RenderList(w, r, SerializeList(characters))
}

func PlayMyCharacter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	character := ctx.Value("character").(*model.Character)

	form := characterService.PlayForm{Character: character}

	characterService.Play(form)
}

func LogoutMyCharacter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	character := ctx.Value("character").(*model.Character)

	form := characterService.LogoutForm{Character: character}

	characterService.Logout(form)
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
