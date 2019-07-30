package main

import (
	"log"
	"net/http"

	"azathot/controller"
	"azathot/router"

	"github.com/gorilla/handlers"
	"github.com/unrolled/render"
)

func main() {
	renderer := render.New()

	playerController := controller.NewPlayer(renderer)
	userController := controller.NewUser(renderer)
	charactersController := controller.NewCharacer(renderer)

	router := router.GetRouter(playerController)
	log.Println("Starting API server in port 1937")
	log.Fatal(http.ListenAndServe(":1937", handlers.CORS(
		handlers.AllowedHeaders(
			[]string{"X-Requested-With", "Content-Type", "Authorization"},
		),
		handlers.AllowedMethods(
			[]string{"GET", "POST", "PATCH", "HEAD", "OPTIONS"},
		),
		handlers.AllowedOrigins(
			[]string{"*"},
		))(router),
	))
}
