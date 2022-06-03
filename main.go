package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Dohvacanje prvog JSON-a

	resp, err := http.Get("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json")
	if err != nil {
		fmt.Println(err)
	}

	bodyOne, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	lige := string(bodyOne)
	fmt.Println(lige)

	//Dohvacanje drugog JSON-a

	res, err := http.Get("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json")
	if err != nil {
		fmt.Println(err)
	}

	bodyTwo, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("")
	}

	ponude := string(bodyTwo)
	fmt.Println(ponude)
}
