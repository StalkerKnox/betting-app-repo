package main

import (
	"betting-app/database"
	"betting-app/handler"
	"betting-app/helper"
	"betting-app/models"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Impementing url variables to read from
var urlLeagues = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json"
var urlOffers = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json"

func main() {

	errLeagues := helper.GetJSON(urlLeagues, &models.Leagues)
	if errLeagues != nil {
		log.Fatal(errLeagues)
	}

	// fmt.Println(models.Leagues)

	fmt.Println("#################################################################################")

	errOffers := helper.GetJSON(urlOffers, &models.Offers)
	if errOffers != nil {
		log.Fatal(errOffers)
	}

	// fmt.Println(models.Offers)

	//Init router

	router := mux.NewRouter()

	// Handling requests

	router.HandleFunc("/leagues", handler.GetLeagues).Methods("GET")
	router.HandleFunc("/offers/{id}", handler.GetOfferbyID).Methods("GET")
	router.HandleFunc("/offers", handler.GetOffers).Methods("GET")
	router.HandleFunc("/offers", handler.AddNewOffer).Methods("POST")

	database.ConnectDB()
	database.InsertToDB()
	database.GetOfferFromDB()
	database.GetOffersFromDB()

	log.Fatal(http.ListenAndServe(":8081", router))
}
