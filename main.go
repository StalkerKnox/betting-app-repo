package main

import (
	"booking-app/controller"
	"booking-app/models"
	"booking-app/structure"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Impementing url variables to read from
var urlLeagues = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json"
var urlOffers = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json"

func main() {

	errLeagues := structure.GetJSON(urlLeagues, &models.Leagues)
	if errLeagues != nil {
		log.Fatal(errLeagues)
	}

	fmt.Println(models.Leagues)

	fmt.Println("#################################################################################")

	errOffers := structure.GetJSON(urlOffers, &models.Offers)
	if errOffers != nil {
		log.Fatal(errOffers)
	}

	fmt.Println(models.Offers)

	//Init router

	router := mux.NewRouter()

	// Handling requests

	router.HandleFunc("/leagues", controller.GetLeagues).Methods("GET")
	router.HandleFunc("/offers/{id}", controller.GetOfferbyID).Methods("GET")
	router.HandleFunc("/offers", controller.GetOffers).Methods("GET")
	router.HandleFunc("/offers", controller.AddNewOffer).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
