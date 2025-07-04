package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"tourist-site/database"
	"tourist-site/routes"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Warning: no .env file loaded (%v)", err)
	}
}

func main() {
	database.InitDB()

	router := routes.RegisterRoutes()

	fs := http.FileServer(http.Dir("../frontend"))
	router.PathPrefix("/").Handler(fs)

	fmt.Println("✅ Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
