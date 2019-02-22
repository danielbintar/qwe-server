package websocket

import (
	"net/http"

	controller "github.com/danielbintar/qwe-server/controller/websocket"
	webController "github.com/danielbintar/qwe-server/controller/web"


	"github.com/go-chi/chi"
)

func Main(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Use(webController.Authenticated)
		hub := controller.HubInstance()
		go hub.Run()
		r.Get("/play", func(w http.ResponseWriter, r *http.Request) { controller.Manage(hub, w, r) })
	})
}
