package models

import (
	"database/sql"
)

type PlaceRaw struct {
	ID              int
	AttractionID    sql.NullString
	AttractionName  sql.NullString
	Category        sql.NullString
	Categories      sql.NullString
	Rating          sql.NullFloat64
	Reviews         sql.NullInt64
	Address         sql.NullString
	City            sql.NullString
	Country         sql.NullString
	Province        sql.NullString
	Zipcode         sql.NullString
	BroaderCategory sql.NullString
	WeightedScore   sql.NullFloat64
	WeightedAverage sql.NullFloat64
	AllCities       sql.NullString
	Description     sql.NullString
	Latitude        sql.NullString
	Longitude       sql.NullString
}

type PlaceJSON struct {
	ID              int     `json:"id"`
	AttractionID    string  `json:"attaction_id"`
	AttractionName  string  `json:"attraction_name"`
	Category        string  `json:"category"`
	Categories      string  `json:"categories"`
	Rating          float64 `json:"rating"`
	Reviews         int     `json:"reviews"`
	Address         string  `json:"address"`
	City            string  `json:"city"`
	Country         string  `json:"country"`
	Province        string  `json:"province"`
	Zipcode         string  `json:"zipcode"`
	BroaderCategory string  `json:"broader_category"`
	WeightedScore   float64 `json:"Weighted_Score"`
	WeightedAverage float64 `json:"Weighted_Average"`
	AllCities       string  `json:"All_Cities"`
	Description     string  `json:"description"`
	Latitude        string  `json:"latitude"`
	Longitude       string  `json:"longitude"`
}

func ConvertToPlaceJSON(p PlaceRaw) PlaceJSON {
	return PlaceJSON{
		ID:              p.ID,
		AttractionID:    nullStringToString(p.AttractionID),
		AttractionName:  nullStringToString(p.AttractionName),
		Category:        nullStringToString(p.Category),
		Categories:      nullStringToString(p.Categories),
		Rating:          nullFloatToFloat(p.Rating),
		Reviews:         nullIntToInt(p.Reviews),
		Address:         nullStringToString(p.Address),
		City:            nullStringToString(p.City),
		Country:         nullStringToString(p.Country),
		Province:        nullStringToString(p.Province),
		Zipcode:         nullStringToString(p.Zipcode),
		BroaderCategory: nullStringToString(p.BroaderCategory),
		WeightedScore:   nullFloatToFloat(p.WeightedScore),
		WeightedAverage: nullFloatToFloat(p.WeightedAverage),
		AllCities:       nullStringToString(p.AllCities),
		Description:     nullStringToString(p.Description),
		Latitude:        nullStringToString(p.Latitude),
		Longitude:       nullStringToString(p.Longitude),
	}
}

func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func nullFloatToFloat(nf sql.NullFloat64) float64 {
	if nf.Valid {
		return nf.Float64
	}
	return 0.0
}

func nullIntToInt(ni sql.NullInt64) int {
	if ni.Valid {
		return int(ni.Int64)
	}
	return 0
}
