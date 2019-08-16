package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

//StatusController with the handler methods relevant to the router
type StatusController interface {
	Healthz(w http.ResponseWriter, req *http.Request)
}

//UserController with the handler methods relevant to the router
type UserController interface {
	Login(w http.ResponseWriter, req *http.Request)
	Signup(w http.ResponseWriter, req *http.Request)
}

type PlayerUsecase interface {
	GetPlayers(w http.ResponseWriter, req *http.Request)
}

//create a new Router
func GetRouter(sc StatusController, uc UserController, pc PlayerUsecase) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", sc.Healthz).Methods("GET").Name("healthz")

	r.HandleFunc("/login", uc.Login).Methods("GET").Name("login")
	r.HandleFunc("/signup", uc.Signup).Methods("GET").Name("signup")

	r.HandleFunc("/players", pc.GetPlayers).Methods("GET").Name("players")

	return r
}
