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

func main() {

	// Impementing url variables to read from
	var urlLeagues = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json"
	var urlOffers = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json"

	var leaguesStruct models.GetLeagueResponse
	var offersSlice models.GetOfferResponse

	errLeagues := helper.GetJSON(urlLeagues, &leaguesStruct)
	if errLeagues != nil {
		log.Fatal(errLeagues)
	}

	errOffers := helper.GetJSON(urlOffers, &offersSlice)
	if errOffers != nil {
		log.Fatal(errOffers)
	}

	//Init router

	router := mux.NewRouter()

	//Database operations
	// database.ConnectDB()
	err := database.InsertOffersIntoDB(offersSlice)
	if err != nil {
		log.Fatal(err)
	}
	err = database.InsertLeaguesIntoDB(leaguesStruct)
	if err != nil {
		log.Fatal(err)
	}

	// Handling requests
	router.HandleFunc("/leagues", handler.GetLeagues).Methods("GET")
	router.HandleFunc("/offers/{id}", handler.GetOfferbyID).Methods("GET")
	router.HandleFunc("/offers", handler.GetOffers).Methods("GET")
	router.HandleFunc("/offers", handler.AddNewOffer).Methods("POST")
	router.HandleFunc("/tickets", handler.AddNewTicket).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}
