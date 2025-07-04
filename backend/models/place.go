package models

type Place struct {
	ID              int     `json:"id"`
	AttractionID    string  `json:"attaction_id"`
	AttractionName  string  `json:"attraction_name"`
	Category        string  `json:"category"`
	Categories      string  `json:"categories"`
	Rating          float32 `json:"rating"`
	Reviews         int     `json:"reviews"`
	Address         string  `json:"address"`
	City            string  `json:"city"`
	Country         string  `json:"country"`
	Province        string  `json:"province"`
	Zipcode         string  `json:"zipcode"`
	BroaderCategory string  `json:"broader_category"`
	WeightedScore   float32 `json:"Weighted_Score"`
	WeightedAverage float32 `json:"Weighted_Average"`
	AllCities       string  `json:"All_Cities"`
	Description     string  `json:"description"`
	Latitude        string  `json:"latitude"`
	Longitude       string  `json:"longitude"`
}
