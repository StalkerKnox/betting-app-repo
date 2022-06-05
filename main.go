package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Liga struct {
	Naziv   string `json:"naziv"`
	Razrade []struct {
		Tipovi []struct {
			Naziv string `json:"naziv"`
		} `json:"tipovi"`
		Ponude []int `json:"ponude"`
	} `json:"razrade"`
}

type Ponuda struct {
	Broj     string    `json:"broj"`
	ID       int       `json:"id"`
	Naziv    string    `json:"naziv"`
	Vrijeme  time.Time `json:"vrijeme"`
	Tecajevi []struct {
		Tecaj float64 `json:"tecaj"`
		Naziv string  `json:"naziv"`
	} `json:"tecajevi"`
	TvKanal       string `json:"tv_kanal,omitempty"`
	ImaStatistiku bool   `json:"ima_statistiku,omitempty"`
}

func main() {
	// Dohvacanje prvog JSON-a

	resp, err := http.Get("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.jso")
	if err != nil {
		fmt.Println(err)

	}

	bodyLige, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var lige []Liga
	jsonErrLige := json.Unmarshal(bodyLige, &lige)
	if jsonErrLige != nil {
		log.Println("Error unmarshalling json data:", err)
	}
	fmt.Println(lige)

	// ligeJSON := string(bodyPonude)
	// fmt.Println(ligeJSON)

	fmt.Print("#######################################################################################################")

	//Dohvacanje drugog JSON-a

	res, err := http.Get("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json")
	if err != nil {
		fmt.Println(err)
	}

	bodyPonude, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var ponude []Ponuda
	jsonErrPonude := json.Unmarshal(bodyPonude, &ponude)
	if jsonErrPonude != nil {
		log.Fatal(jsonErrPonude)
	}
	fmt.Print(ponude)

	// ponudeJSON := string(bodyPonude)
	// fmt.Println(ponudeJSON)

}
