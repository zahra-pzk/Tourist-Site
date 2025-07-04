package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"tourist-site/database"
	"tourist-site/models"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getGoogleConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	config := getGoogleConfig()
	url := config.AuthCodeURL("random_state")
	log.Println("🔁 Redirecting to:", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	config := getGoogleConfig()

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "No code in request", http.StatusBadRequest)
		return
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	client := config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Picture string `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to parse user info", http.StatusInternalServerError)
		return
	}

	db := database.DB

	var existingUser models.User
	query := "SELECT id, google_id, name, email, picture FROM users WHERE email = $1 LIMIT 1"
	err = db.QueryRow(query, userInfo.Email).Scan(&existingUser.ID, &existingUser.GoogleID, &existingUser.Name, &existingUser.Email, &existingUser.Picture)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			insertQuery := `INSERT INTO users (google_id, name, email, picture) VALUES ($1, $2, $3, $4)`
			_, insertErr := db.Exec(insertQuery, userInfo.ID, userInfo.Name, userInfo.Email, userInfo.Picture)
			if insertErr != nil {
				log.Println("❌ Error creating user:", insertErr)
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}
			log.Println("✅ New user created:", userInfo.Email)
		} else {
			log.Println("❌ Error querying user:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("ℹ️ User already exists:", existingUser.Email)
	}

	http.Redirect(w, r, "/index.html", http.StatusSeeOther)
}
