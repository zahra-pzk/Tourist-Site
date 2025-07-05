package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tourist-site/database"
	"tourist-site/models"
)

func GetAllPlaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	country := r.URL.Query().Get("country")
	query := `
		SELECT id, attaction_id, attraction_name, category, categories, rating, reviews, address, city, country,
		province, zipcode, broader_category, weighted_score, weighted_average, all_cities, description, latitude, longitude
		FROM tourist_places`
	var rows *sql.Rows
	var err error

	if country != "" {
		query += " WHERE LOWER(country) = LOWER($1)"
		rows, err = database.DB.Query(query, country)
	} else {
		rows, err = database.DB.Query(query)
	}
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		log.Println("❌ Query error:", err)
		return
	}
	defer rows.Close()

	var results []models.PlaceJSON
	for rows.Next() {
		var p models.PlaceRaw
		err := rows.Scan(
			&p.ID,
			&p.AttractionID, &p.AttractionName, &p.Category, &p.Categories,
			&p.Rating, &p.Reviews, &p.Address, &p.City, &p.Country,
			&p.Province, &p.Zipcode, &p.BroaderCategory, &p.WeightedScore,
			&p.WeightedAverage, &p.AllCities, &p.Description, &p.Latitude, &p.Longitude,
		)
		if err != nil {
			http.Error(w, "Error scanning place", http.StatusInternalServerError)
			log.Println("❌ Error scanning place:", err)
			return
		}
		results = append(results, models.ConvertToPlaceJSON(p))
	}
	json.NewEncoder(w).Encode(results)
}

func GetPlaceByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	row := database.DB.QueryRow(`
		SELECT id, attaction_id, attraction_name, category, categories, rating, reviews, address, city, country,
		province, zipcode, broader_category, weighted_score, weighted_average, all_cities, description, latitude, longitude
		FROM tourist_places WHERE id = $1`, id)

	var p models.PlaceRaw
	err := row.Scan(
		&p.ID,
		&p.AttractionID, &p.AttractionName, &p.Category, &p.Categories,
		&p.Rating, &p.Reviews, &p.Address, &p.City, &p.Country,
		&p.Province, &p.Zipcode, &p.BroaderCategory, &p.WeightedScore,
		&p.WeightedAverage, &p.AllCities, &p.Description, &p.Latitude, &p.Longitude,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Place not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
			log.Println("❌ DB error:", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ConvertToPlaceJSON(p))
}