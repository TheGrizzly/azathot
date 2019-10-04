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

type PlayerController interface {
	GetPlayers(w http.ResponseWriter, req *http.Request)
	GetPlayer(w http.ResponseWriter, req *http.Request)
	PostPlayer(w http.ResponseWriter, req *http.Request)
	PatchPlayer(w http.ResponseWriter, req *http.Request)
	DeletePlayerById(w http.ResponseWriter, req *http.Request)
}

type JWTMiddelware interface {
	ValidateJWT(handler http.HandlerFunc) http.HandlerFunc
}

//create a new Router
func GetRouter(sc StatusController, uc UserController, pc PlayerController, jm JWTMiddelware) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", sc.Healthz).Methods("GET").Name("healthz")

	r.HandleFunc("/login", uc.Login).Methods("POST").Name("login")
	r.HandleFunc("/signup", uc.Signup).Methods("POST").Name("signup")

	r.HandleFunc("/players", pc.GetPlayers).Methods("GET").Name("getPlayers")
	r.HandleFunc("/players/{player_id}", pc.GetPlayer).Methods("GET").Name("getPlayer")
	r.HandleFunc("/admin/players", jm.ValidateJWT(pc.PostPlayer)).Methods("POST").Name("postPlayers")
	r.HandleFunc("/admin/players/{player_id}", jm.ValidateJWT(pc.PatchPlayer)).Methods("PATCH").Name("patchPlayers")
	r.HandleFunc("/admin/players/{player_id}", jm.ValidateJWT(pc.DeletePlayerById)).Methods("DELETE").Name("deletePlayer")

	return r
}
