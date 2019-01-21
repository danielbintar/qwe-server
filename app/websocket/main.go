package websocket

import (
	"net/http"

	controller "github.com/danielbintar/qwe-server/websocket_controller"

	"github.com/go-chi/chi"
)

func Main(r *chi.Mux) {
	chatHub := controller.NewHub()
	go chatHub.Run()
	r.Get("/chat", func(w http.ResponseWriter, r *http.Request) { controller.ManageChat(chatHub, w, r) })

	moveHub := controller.NewHub()
	go moveHub.Run()
	r.Get("/move", func(w http.ResponseWriter, r *http.Request) { controller.ManageMove(moveHub, w, r) })
}
