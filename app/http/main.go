package main

import (
	"fmt"
	"net/http"

	"github.com/danielbintar/qwe-server/app/websocket"
	"github.com/danielbintar/qwe-server/controller"

	"github.com/go-chi/chi"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	r := chi.NewRouter()

	hub := websocket.NewHub()
	go hub.Run()

	r.Post("/users/sign_in", controller.Login)
	r.Get("/chat", func(w http.ResponseWriter, r *http.Request) { websocket.ServeWs(hub, w, r) })

	fmt.Println("listen to 3333")
	http.ListenAndServe(":3333", r)
}
