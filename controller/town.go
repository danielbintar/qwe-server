package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"
	"github.com/danielbintar/qwe-server/service/auth"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Town(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		townId, err := strconv.Atoi(chi.URLParam(r, "townId"))
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		town := repository.FindTown(townId)
		if town.Name == "" {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "town", town)


		form := auth.LoginForm{}

		switch r.Method {
		case http.MethodGet:
			form.Username = r.URL.Query().Get("username")
			form.Password = r.URL.Query().Get("password")
		case http.MethodPost:
			err = json.NewDecoder(r.Body).Decode(&form)
			if err != nil {
				http.Error(w, http.StatusText(401), 401)
				return
			}
		}

		userI, errors := auth.Login(form)
		if errors != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		byteData, _ := json.Marshal(userI)
		var user *model.User
		json.Unmarshal(byteData, &user)
		ctx = context.WithValue(ctx, "user", user)


		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FindTown(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	town, _ := ctx.Value("town").(*model.Town)
	town.Users = repository.GetTownUsers(town.Id)

	render.Render(w, r, town)
}

func EnterTown(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	town, _ := ctx.Value("town").(*model.Town)

	currentUser, _ := ctx.Value("user").(*model.User)
	repository.SetTownUser(town.Id, currentUser.Id, town.Position.X, town.Position.Y)

	town.Users = repository.GetTownUsers(town.Id)
	render.Render(w, r, town)
}
