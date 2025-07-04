package main

import (
    "fmt"
    "log"
    "net/http"
    "tourist-site/database"
	"tourist-site/routes"

    //"github.com/gorilla/mux"
)

func main() {
database.InitDB()
r := routes.RegisterRoutes()
fmt.Println("✅ Server running on http://localhost:8080")
log.Fatal(http.ListenAndServe(":8080", r))
}