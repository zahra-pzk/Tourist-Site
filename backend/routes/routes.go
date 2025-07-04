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

    staticFileDirectory := http.Dir("../frontend")
    staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
    r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

    return r
}
