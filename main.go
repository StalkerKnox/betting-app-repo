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
	Type []Type `json:"tipovi"`
	Name []int  `json:"ponude"`
}

type Type struct {
	Name string `json:"Naziv"`
}

// Creating GetOfferResponse Struct

type GetOfferResponse struct {
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

func main() {
	// Getting First JSON

	resp, err := http.Get("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	leaugesBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var leagues GetLeagueResponse
	jsonErrLeagues := json.Unmarshal(leaugesBody, &leagues)
	if jsonErrLeagues != nil {
		log.Fatal("Error unmarshalling json data:", err)
	}
	fmt.Println(leagues)

	fmt.Println("#######################################################################################################")

	//Getting second JSON

	res, err := http.Get("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json")
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	offersBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var offers []GetOfferResponse
	jsonErrPonude := json.Unmarshal(offersBody, &offers)
	if jsonErrPonude != nil {
		log.Fatal(jsonErrPonude)
	}
	fmt.Print(offers)
}
