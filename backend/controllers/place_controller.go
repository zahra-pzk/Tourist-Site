package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	//"strconv"

	"tourist-site/database"
	"tourist-site/models"

	"github.com/gorilla/mux"
)

func GetAllPlaces(w http.ResponseWriter, r *http.Request) {
rows, err := database.DB.Query("SELECT id, attaction_id, attraction_name, category, categories, rating, reviews, address, city, country, province, zipcode, broader_category, Weighted_Score, Weighted_Average, All_Cities, description, latitude, longitude FROM tourist_places")
if err != nil {
http.Error(w, "Database error", http.StatusInternalServerError)
return
}
defer rows.Close()

var places []models.Place
for rows.Next() {
	var p models.Place
	err := rows.Scan(&p.ID, &p.AttractionID, &p.AttractionName, &p.Category, &p.Categories, &p.Rating, &p.Reviews, &p.Address, &p.City, &p.Country, &p.Province, &p.Zipcode, &p.BroaderCategory, &p.WeightedScore, &p.WeightedAverage, &p.AllCities, &p.Description, &p.Latitude, &p.Longitude)
	if err != nil {
		http.Error(w, "Scan error", http.StatusInternalServerError)
		return
	}
	places = append(places, p)
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(places)
}

func GetPlaceByID(w http.ResponseWriter, r *http.Request) {
id := mux.Vars(r)["id"]
row := database.DB.QueryRow("SELECT id, attaction_id, attraction_name, category, categories, rating, reviews, address, city, country, province, zipcode, broader_category, Weighted_Score, Weighted_Average, All_Cities, description, latitude, longitude FROM tourist_places WHERE id = $1", id)

var p models.Place
err := row.Scan(&p.ID, &p.AttractionID, &p.AttractionName, &p.Category, &p.Categories, &p.Rating, &p.Reviews, &p.Address, &p.City, &p.Country, &p.Province, &p.Zipcode, &p.BroaderCategory, &p.WeightedScore, &p.WeightedAverage, &p.AllCities, &p.Description, &p.Latitude, &p.Longitude)
if err != nil {
	if err == sql.ErrNoRows {
		http.Error(w, "Place not found", http.StatusNotFound)
	} else {
		http.Error(w, "Database error", http.StatusInternalServerError)
	}
	return
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(p)
}