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
	Number        string    `json:"broj" db:"number" validate:"required"`
	TVchannel     string    `json:"tv_kanal" db:"tv_channel"`
	ID            int       `json:"id" db:"offer_id"`
	Title         string    `json:"naziv" db:"title" validate:"required"`
	HasStatistics bool      `json:"ima_statistiku" db:"has_statistics" validate:"required"`
	Time          time.Time `json:"vrijeme" db:"time" validate:"required"`
	Rates         []Rate    `json:"tecajevi" validate:"required"`
}

type Rate struct {
	OfferID int     `db:"offer_id"`
	Rate    float64 `json:"tecaj" db:"rate" validate:"required"`
	Name    string  `json:"naziv" db:"name" validate:"required"`
}

// Defining structure variables to store parsed JSON
var Leagues GetLeagueResponse
var Offers GetOfferResponse
var OneOffer Offer
