package main

import (
    "fmt"
    "log"
    "net/http"
    "tourist-site/database"

    "github.com/gorilla/mux"
)

func main() {
    database.InitDB()

    r := mux.NewRouter()

    fmt.Println("🚀 Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
