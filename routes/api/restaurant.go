package api

import (
	"database/sql"
	"encoding/json"
	"github.com/antonybudianto/resto-api/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const radius = 5
const top = 11

// Handler for route
type Handler struct {
	Router *mux.Router
	DB     *sql.DB
}

func (handler *Handler) getNearestRestaurants(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	latKeys, _ := query["lat"]
	lngKeys, _ := query["lng"]
	lat := latKeys[0]
	lng := lngKeys[0]

	log.Printf("GET nearest resto lat:%s lng:%s", lat, lng)
	restaurants, err := model.GetRestaurants(handler.DB, 0, top)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, restaurants)
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
}
