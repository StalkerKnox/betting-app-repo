package main

import (
	"betting-app/database"
	"betting-app/handler"
	"betting-app/helper"
	"betting-app/models"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Impementing url variables to read from
var urlLeagues = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json"
var urlOffers = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json"

func main() {

	errLeagues := helper.GetJSON(urlLeagues, &models.LeaguesStruct)
	if errLeagues != nil {
		log.Fatal(errLeagues)
	}

	errOffers := helper.GetJSON(urlOffers, &models.Offers)
	if errOffers != nil {
		log.Fatal(errOffers)
	}

	//Init router

	router := mux.NewRouter()

	//Database operations
	database.ConnectDB()
	database.InsertOffersIntoDB()
	database.InsertLeaguesIntoDB()

	// Handling requests
	router.HandleFunc("/leagues", handler.GetLeagues).Methods("GET")
	router.HandleFunc("/offers/{id}", handler.GetOfferbyID).Methods("GET")
	router.HandleFunc("/offers", handler.GetOffers).Methods("GET")
	router.HandleFunc("/offers", handler.AddNewOffer).Methods("POST")
	router.HandleFunc("/tickets", handler.AddNewTicket).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}
