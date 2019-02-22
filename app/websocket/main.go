package websocket

import (
	"net/http"

	controller "github.com/danielbintar/qwe-server/websocket_controller"
	webController "github.com/danielbintar/qwe-server/controller/web"


	"github.com/go-chi/chi"
)

func Main(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Use(webController.Authenticated)
		chatHub := controller.ChatHubInstance()
		go chatHub.Run()
		r.Get("/chat", func(w http.ResponseWriter, r *http.Request) { controller.ManageChat(chatHub, w, r) })

		moveHub := controller.MoveHubInstance()
		go moveHub.Run()
		r.Get("/move", func(w http.ResponseWriter, r *http.Request) { controller.ManageMove(moveHub, w, r) })
	})
}
