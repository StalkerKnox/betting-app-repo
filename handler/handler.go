package handler

import (
	"betting-app/database"
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
	err := database.GetLeaguesFromDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(models.GetLeaguesFromDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GET offers / implemented just for checking POST method
func GetOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := database.GetOffersFromDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(models.GetOffersFromDB)
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
	dataErr := database.GetOfferFromDB(id)
	if dataErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Offer does not exists"))
		return
	}
	json.NewEncoder(w).Encode(models.OfferFromDB)

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
	// for _, rate := range offer.Rates {
	// 	rate.OfferID = offer.ID
	// }
	insertErr := database.InsertOfferIntoDB(offer)
	if insertErr != nil {
		return
	}

	json.NewEncoder(w).Encode(offer)
}
