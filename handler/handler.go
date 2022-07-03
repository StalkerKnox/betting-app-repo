package handler

import (
	"betting-app/database"
	"betting-app/models"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// GET leagues
func GetLeagues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	getLeaguesFromDB, err := database.GetLeaguesFromDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(getLeaguesFromDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GET offers / implemented just for checking POST method
func GetOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	getOffersFromDB, err := database.GetOffersFromDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(getOffersFromDB)
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
	offerFromDB, dataErr := database.GetOfferFromDB(id)
	if dataErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Offer does not exists"))
		return
	}
	json.NewEncoder(w).Encode(offerFromDB)

}

// ADD new offer (POST method)
func AddNewOffer(w http.ResponseWriter, r *http.Request) {

	var offer models.Offer
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

// Simulate ticket (POST method)
func AddNewTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.TikcetDesign
	json.NewDecoder(r.Body).Decode(&ticket)
	var err error
	ticket.RemainingBalance, err = database.GetBalanceFromDB(ticket)
	if err != nil {
		log.Fatal(err)
	}
	if ticket.PaymentAmount > ticket.RemainingBalance {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sorry, but your account balance is not sufficient for this payment amaount."))
		return
	}
	playedTicket, _ := database.GetRatesFromDB(ticket)
	if playedTicket.PrizeMoney > 10000 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sorry, but the prize money exceeds 10,000 "))
		return
	}

	_ = database.UpdateBalance(*playedTicket)
	playedTicket, err = database.InsertTicketIntoDB(*playedTicket)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(playedTicket)
	fmt.Println(playedTicket)

}
