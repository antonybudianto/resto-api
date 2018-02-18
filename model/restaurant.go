package model

// Restaurant data
type Restaurant struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Slug      string  `json:"slug"`
	CuisineID string  `json:"cuisineId"`
	CountryID string  `json:"countryId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	Rating    float32 `json:"rating"`
}
