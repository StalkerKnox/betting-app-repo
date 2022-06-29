package handler

import (
	"betting-app/database"
	"betting-app/models"
	"encoding/json"
	"fmt"
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
	insertErr := database.InsertOfferIntoDB(offer)
	if insertErr != nil {
		return
	}

	json.NewEncoder(w).Encode(offer)
}

// Simulate ticket
func AddNewTicket(w http.ResponseWriter, r *http.Request) {

	json.NewDecoder(r.Body).Decode(&models.Ticket)
	_ = database.GetBalanceFromDB()
	if models.Ticket.PaymentAmount > models.Ticket.RemainingBalance {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sorry, but you are broke. Come back when you are rich"))
		return
	}
	_ = database.GetRatesFromDB()
	if models.Ticket.PrizeMoney > 10000 {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Sorry, but we don't accept values over 10k"))
	}

	_ = database.UpdateBalance()
	json.NewEncoder(w).Encode(models.Ticket)
	fmt.Println(models.Ticket)

}
