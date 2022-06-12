package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Impementing url variables to read from

var urlLeagues = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json"
var urlOffers = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json"

// Creating GetLeagueResponse Struct

type GetLeagueResponse struct {
	League []League `json:"lige"`
}

type League struct {
	Title       string        `json:"naziv"`
	Elaboration []Elaboration `json:"razrade"`
}

type Elaboration struct {
	Type   []Type `json:"tipovi"`
	Offers []int  `json:"ponude"`
}

type Type struct {
	Name string `json:"naziv"`
}

//Creating GetOfferResponse Struct

type GetOfferResponse []Offer

type Offer struct {
	Number        string    `json:"broj"`
	TVchannel     string    `json:"tv_kanal"`
	ID            int       `json:"id"`
	Title         string    `json:"naziv"`
	HasStatistics bool      `json:"ima_statistiku"`
	Time          time.Time `json:"vrijeme"`
	Rate          []Rate    `json:"tecajevi"`
}

type Rate struct {
	Rates float64 `json:"tecaj"`
	Name  string  `json:"naziv"`
}

// Defining structure variables to store parsed JSON

var leagues GetLeagueResponse
var offers GetOfferResponse

// Parsing JSON

func getJSON(URL string, structure interface{}) {
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(resBody, structure)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
}

// GET leagues

func getLeagues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leagues)
}

// GET offers / implemented just for checking POST method

func getOffers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(offers)
}

// GET offers by ID

func getOfferbyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range offers {
		id := strconv.Itoa(item.ID)
		if id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

// ADD new offer (POST method)

func addNewOffer(w http.ResponseWriter, r *http.Request) {
	var offer Offer
	_ = json.NewDecoder(r.Body).Decode(&offer)
	offer.ID = rand.Intn(10000000)
	offers = append(offers, offer)
	json.NewEncoder(w).Encode(offer)
}

func main() {

	getJSON(urlLeagues, &leagues)
	fmt.Println(leagues)

	getJSON(urlOffers, &offers)
	fmt.Println(offers)

	//Init router

	router := mux.NewRouter()

	// Handling requests

	router.HandleFunc("/leagues", getLeagues).Methods("GET")
	router.HandleFunc("/offers/{id}", getOfferbyID).Methods("GET")
	router.HandleFunc("/offers", getOffers).Methods("GET")
	router.HandleFunc("/offers", addNewOffer).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
