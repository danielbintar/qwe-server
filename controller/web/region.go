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

func Region(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u64, err := strconv.ParseUint(chi.URLParam(r, "regionID"), 10, 32)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		regionID := uint(u64)

		region := repository.FindRegion(regionID)
		if region.Name == "" {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		ctx := context.WithValue(r.Context(), "region", region)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func FindRegion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	region, _ := ctx.Value("region").(*model.Region)
	region.CharactersPosition = repository.GetRegionCharactersPosition(region.ID)

	serializer := region.Serialize()
	serializer.Monsters = repository.AllSpawnMonster(region.ID)

	render.Render(w, r, serializer)
}
