package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("SSL_MODE"),
    )

    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Failed to ping DB:", err)
    }

    fmt.Println("✅ Connected to PostgreSQL!")
}
