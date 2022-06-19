package models

import "time"

// Creating GetLeagueResponse Struct
type GetLeagueResponse struct {
	Leagues []League `json:"lige"`
}

type League struct {
	Title        string        `json:"naziv"`
	Elaborations []Elaboration `json:"razrade"`
}

type Elaboration struct {
	Types  []Type `json:"tipovi"`
	Offers []int  `json:"ponude"`
}

type Type struct {
	Name string `json:"naziv"`
}

//Creating GetOfferResponse Struct
type GetOfferResponse []Offer

type Offer struct {
	Number        string    `json:"broj" validate:"required"`
	TVchannel     string    `json:"tv_kanal"`
	ID            int       `json:"id"`
	Title         string    `json:"naziv" validate:"required"`
	HasStatistics bool      `json:"ima_statistiku" validate:"required"`
	Time          time.Time `json:"vrijeme" validate:"required"`
	Rates         []Rate    `json:"tecajevi" validate:"required"`
}

type Rate struct {
	Rate float64 `json:"tecaj" validate:"required"`
	Name string  `json:"naziv" validate:"required"`
}

// Defining structure variables to store parsed JSON
var Leagues GetLeagueResponse
var Offers GetOfferResponse
var OneOffer Offer
