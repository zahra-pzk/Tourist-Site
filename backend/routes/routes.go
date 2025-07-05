package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"tourist-site/controllers"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/places", controllers.GetAllPlaces).Methods("GET")
	r.HandleFunc("/auth/google/login", controllers.GoogleLogin)
	r.HandleFunc("/auth/google/callback", controllers.GoogleCallback)

	fs := http.FileServer(http.Dir("../frontend"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	return r
}
