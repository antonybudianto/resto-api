package model

// Book for Booking table
type Book struct {
	ID           int `json:"id"`
	UserID       int `json:"userId"`
	RestaurantID int `json:"restaurantId"`
	BookDatetime int `json:"bookDatetime"`
	TotalPeople  int `json:"totalPeople"`
}
