package model

import (
	"database/sql"
	"fmt"
)

const nearestKm = 5

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
	Distance  float64 `json:"distance"`
}

// GetRestaurants : get restaurants
// source: https://developers.google.com/maps/articles/phpsqlsearch_v3
func GetRestaurants(db *sql.DB, start, count int, lat, lng float64) ([]Restaurant, error) {
	statement := fmt.Sprintf(`
	SELECT id, name, slug, cuisine_id, country_id, lat, lng, address, rating,
	( 6371 * acos( cos( radians(%f) ) *
	cos( radians( lat ) ) *
	cos( radians( lng ) - radians(%f) ) +
	sin( radians(%f) ) * sin( radians( lat ) ) ) ) AS distance
	FROM restaurants
	HAVING distance < %d
	ORDER BY distance
	LIMIT %d OFFSET %d`, lat, lng, lat, nearestKm, count, start)
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
			&r.Rating,
			&r.Distance); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, r)
	}

	return restaurants, nil
}
