package websocket

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Main(r *chi.Mux) {
	chatHub := NewHub()
	go chatHub.Run()
	r.Get("/chat", func(w http.ResponseWriter, r *http.Request) { ManageChat(chatHub, w, r) })
}
