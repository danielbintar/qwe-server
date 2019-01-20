package controller

import (
	"context"
	"net/http"
	"strconv"

	townConfig "github.com/danielbintar/qwe-server/config/town"
	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"

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

		town  := townConfig.Find(townId)
		if town.Name == "" {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "town", town)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FindTown(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	town, _ := ctx.Value("town").(*model.Town)
	town.Users = repository.GetTownUsers(town.Id)

	render.Render(w, r, town)
}
