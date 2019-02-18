package controller

import (
	"context"
	"net/http"

	"github.com/danielbintar/qwe-server/db"
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"

	"github.com/go-chi/render"
)

func MyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user := &model.User{ID: ctx.Value("jwt").(*model.Jwt).UserID}
		db.DB().Where(&user).First(&user)

		ctx = context.WithValue(ctx, "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetMyUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := ctx.Value("user").(*model.User)

	render.Render(w, r, currentUser.Serialize())
}

func GetMyUserCurrentCharacter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := ctx.Value("user").(*model.User)
	currentCharacterID := repository.GetPlayingCharacter(currentUser.ID)

	currentCharacter := &model.CurrentCharacter{
		CharacterID: currentCharacterID,
	}

	render.Render(w, r, currentCharacter)
}
