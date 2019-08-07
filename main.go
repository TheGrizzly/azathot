package main

import (
	"log"
	"net/http"

	"azathot/config"
	"azathot/controller"
	"azathot/router"
	"azathot/service/crypt"
	"azathot/service/database"
	"azathot/usecase"

	"github.com/gorilla/handlers"
	"github.com/unrolled/render"
)

func main() {
	renderer := render.New()

	appConfig, err := config.LoadFromConfigFile()
	if err != nil {
		log.Fatal("error retreiving configuration: ", err)
	}

	db, err := database.New(appConfig)
	if err != nil {
		log.Fatal("error creating db service")
	}

	healthChecker := controller.HealthChecker{
		db,
	}

	if err = healthChecker.IsHealthy(); err != nil {
		log.Fatal("error checking deps of healthiness: ", err)
	}

	cryptService := crypt.New(appConfig)

	userUsecase := usecase.NewUser(db, cryptService)

	userController := controller.NewUser(userUsecase, renderer)
	statusController := controller.NewStatus(renderer, healthChecker)
	router := router.GetRouter(statusController, userController)
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
