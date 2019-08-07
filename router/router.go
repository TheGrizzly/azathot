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

//create a new Router
func GetRouter(sc StatusController, uc UserController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", sc.Healthz).Methods("GET").Name("healthz")

	r.HandleFunc("/login", uc.Login).Methods("POST").Name("login")
	r.HandleFunc("/signup", uc.Signup).Methods("POST").Name("signup")

	return r
}
