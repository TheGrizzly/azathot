package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

//BienController with the handler methods relevant to the router
type PlayerController interface {
	Healthz(w http.ResponseWriter, req *http.Request)
}

//create a new Router
func GetRouter(pc PlayerController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", pc.Healthz).Methods("GET").Name("healthz")
	return r
}
