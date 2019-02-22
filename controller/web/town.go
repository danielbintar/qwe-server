package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/danielbintar/qwe-server/model"
	"github.com/danielbintar/qwe-server/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Town(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u64, err := strconv.ParseUint(chi.URLParam(r, "townID"), 10, 32)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		townID := uint(u64)

		town := repository.FindTown(townID)
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
	town.CharactersPosition = repository.GetTownCharactersPosition(town.ID)

	render.Render(w, r, town)
}
