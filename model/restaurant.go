package model

import (
	"database/sql"
	"fmt"
)

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

// GetRestaurants : get restaurants
func GetRestaurants(db *sql.DB, start, count int) ([]Restaurant, error) {
	statement := fmt.Sprintf("SELECT id, name, slug, cuisine_id, country_id, lat, lng, address, rating FROM restaurants LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	restaurants := []Restaurant{}

	for rows.Next() {
		var r Restaurant
		if err := rows.Scan(&r.ID, &r.Name, &r.Slug,
			&r.CuisineID,
			&r.CountryID,
			&r.Latitude,
			&r.Longitude,
			&r.Address,
			&r.Rating); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, r)
	}

	return restaurants, nil
}
