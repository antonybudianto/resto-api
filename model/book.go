package model

import (
	"database/sql"
	"fmt"
)

// Book for Booking table
type Book struct {
	ID           int    `json:"id"`
	UserID       int    `json:"userId"`
	RestaurantID int    `json:"restaurantId"`
	BookDatetime string `json:"bookDatetime"`
	TotalPeople  int    `json:"totalPeople"`
}

// CreateBook - Create restaurant booking
func (b *Book) CreateBook(db *sql.DB) (int, error) {
	statement := fmt.Sprintf(`
		INSERT INTO books(user_id, restaurant_id, book_datetime, total_people)
		VALUES(%d, %d, '%s', %d)`,
		b.UserID, b.RestaurantID, b.BookDatetime, b.TotalPeople)
	_, err := db.Exec(statement)

	if err != nil {
		return -1, err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&b.ID)

	if err != nil {
		return -1, err
	}

	return b.ID, nil
}
