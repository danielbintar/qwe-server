package main

import (
	"fmt"
	"net/http"

	"github.com/danielbintar/qwe-server/app/websocket"
	"github.com/danielbintar/qwe-server/controller"
	"github.com/danielbintar/qwe-server/config"

	"github.com/go-chi/chi"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.Instance()
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/sign_in", controller.Login)
		r.Post("/sign_up", controller.CreateUser)
	})

	r.Route("/towns", func(r chi.Router) {
		r.Route("/{townId}", func(r chi.Router) {
			r.Use(controller.Authenticated)
			r.Use(controller.Town)
			r.Get("/", controller.FindTown)
			r.Post("/enter", controller.EnterTown)
		})
	})

	websocket.Main(r)

	fmt.Println("listen to 3333")
	http.ListenAndServe(":3333", r)
}
