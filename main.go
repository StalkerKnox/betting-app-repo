package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

// type GetOfferResponse []Offers

type Offer struct {
	Number        string    `json:"broj"`
	TVchannel     string    `json:"tv_kanal"`
	ID            int       `json:"id"`
	Title         string    `json:"name"`
	HasStatistics bool      `json:"ima_statistiku"`
	Time          time.Time `json:"vrijeme"`
	Rate          []Rate    `json:"tecajevi"`
}

type Rate struct {
	Rates float64 `json:"tecaj"`
	Name  string  `json:"naziv"`
}

var urlLeagues = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json"
var urlOffers = "https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json"

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

func main() {

	var leagues GetLeagueResponse
	var offers GetOfferResponse

	getJSON(urlLeagues, &leagues)
	fmt.Println(leagues)

	getJSON(urlOffers, &offers)
	fmt.Println(offers)
}
