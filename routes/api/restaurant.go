package api

import (
	"database/sql"
	"encoding/json"
	"github.com/antonybudianto/resto-api/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const radius = 5
const top = 11

// Handler for route
type Handler struct {
	Router *mux.Router
	DB     *sql.DB
}

// RestaurantLocationResponse - Custom restaurant location
type RestaurantLocationResponse struct {
	Address   string  `json:"address"`
	CountryID string  `json:"countryId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// RestaurantResponse - Custom restaurant response
type RestaurantResponse struct {
	ID        int                        `json:"id"`
	Name      string                     `json:"name"`
	Slug      string                     `json:"slug"`
	Location  RestaurantLocationResponse `json:"location"`
	CuisineID string                     `json:"cuisineId"`
	Rating    float32                    `json:"rating"`
	Distance  float64                    `json:"distance"`
}

// DataResponse - Restaurant outer response
type DataResponse struct {
	Data []RestaurantResponse `json:"data"`
}

func (handler *Handler) getNearestRestaurants(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	latKeys, _ := query["lat"]
	lngKeys, _ := query["lng"]
	lat, _ := strconv.ParseFloat(latKeys[0], 64)
	lng, _ := strconv.ParseFloat(lngKeys[0], 64)

	log.Printf("GET nearest resto lat:%f lng:%f", lat, lng)
	restaurants, err := model.GetRestaurants(handler.DB, 0, top, lat, lng)
	restoPayload := []RestaurantResponse{}

	for i := 0; i < len(restaurants); i++ {
		resto := restaurants[i]
		restoLocationPayload := RestaurantLocationResponse{
			resto.Address,
			resto.CountryID,
			resto.Latitude,
			resto.Longitude,
		}
		payload := RestaurantResponse{
			resto.ID,
			resto.Name,
			resto.Slug,
			restoLocationPayload,
			resto.CuisineID,
			resto.Rating,
			resto.Distance,
		}
		restoPayload = append(restoPayload, payload)
	}

	dataPayload := DataResponse{restoPayload}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, dataPayload)
}

func (handler *Handler) createBooking(w http.ResponseWriter, r *http.Request) {
	var b model.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&b); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// TODO: User id should be from cookie/token. For demo purpose, I used dummy ID user.
	b.UserID = 1

	id, err := b.CreateBook(handler.DB)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Booking %d created", id)
	respondWithJSON(w, http.StatusCreated, b)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// InitializeRoutes for restaurant endpoints
func (handler *Handler) InitializeRoutes() {
	handler.Router.HandleFunc("/restaurants", handler.getNearestRestaurants).Methods("GET")
	handler.Router.HandleFunc("/books", handler.createBooking).Methods("POST")
}
