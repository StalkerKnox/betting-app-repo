package controller

import (
	"betting-app/models"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// GET leagues
func GetLeagues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(models.Leagues)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GET offers / implemented just for checking POST method
func GetOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Offers)
}

// GET offers by ID
func GetOfferbyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	for _, item := range models.Offers {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Offer does not exists"))
}

// ADD new offer (POST method)
func AddNewOffer(w http.ResponseWriter, r *http.Request) {
	offer := models.OneOffer
	json.NewDecoder(r.Body).Decode(&offer)
	validate := validator.New()
	err := validate.Struct(offer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad input by user"))
		return
	}
	offer.ID = rand.Intn(10000000)
	models.Offers = append(models.Offers, offer)
	json.NewEncoder(w).Encode(offer)
}
