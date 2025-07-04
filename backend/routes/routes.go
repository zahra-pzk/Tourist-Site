package routes

import (
	"net/http"
	"tourist-site/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
r := mux.NewRouter()
r.HandleFunc("/places", controllers.GetAllPlaces).Methods("GET")
r.HandleFunc("/places/{id}", controllers.GetPlaceByID).Methods("GET")
r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("Tourist API is running"))
})
return r
}