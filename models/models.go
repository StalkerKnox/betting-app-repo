package models

// Creating GetLeagueResponse Struct
type GetLeagueResponse struct {
	Leagues []League `json:"lige"`
}

var LeaguesStructDB GetLeagueResponse
var LeaguesDB []League
var LeagueDB League
var ElaborationDB Elaboration
var TypeDB Type
var LeaguesStructDBDB GetLeagueResponse
var Pomagac Elaboration

type League struct {
	Title        string        `json:"naziv" db:"title"`
	Elaborations []Elaboration `json:"razrade"`
	ID           int           `json:"id" db:"id"`
}

type Elaboration struct {
	Types    []Type  `json:"tipovi" `
	Offers   []int64 `json:"ponude"`
	LeagueID int64   `db:"league_id"`
	ID       int64   `db:"elaboration_id"`
}

type Type struct {
	Name          string `json:"naziv" db:"name"`
	ElaborationID int64  `json:"elaboration_id" db:"elaboration_id"`
	ID            int64  `json:"tip_id" db:"type_id"`
}

//Creating GetOfferResponse Struct
type GetOfferResponse []Offer

type Offer struct {
	Number        string `json:"broj" db:"number" validate:"required"`
	TVchannel     string `json:"tv_kanal" db:"tv_channel"`
	ID            int    `json:"id" db:"offer_id"`
	Title         string `json:"naziv" db:"title" validate:"required"`
	HasStatistics bool   `json:"ima_statistiku" db:"has_statistics" validate:"required"`
	Time          string `json:"vrijeme" db:"time" validate:"required"`
	Rates         []Rate `json:"tecajevi" validate:"required"`
}

type Rate struct {
	OfferID int     `json:"ponude_id" db:"offer_id"`
	Rate    float64 `json:"tecaj" db:"rate" validate:"required"`
	Name    string  `json:"naziv" db:"name" validate:"required"`
}

// Defining structure variables to store parsed JSON
var LeaguesStruct GetLeagueResponse
var Offers GetOfferResponse
var OneOffer Offer

//DATABASE

// defining variables to store OFFER BY ID
var OfferFromDB Offer
var RateFromDB Rate

//defining variables to store OFFERS

type OffersFromDB []Offer

var SingleOffer Offer
var IterationStorage OffersFromDB

var SingleRate Rate
var GetOffersFromDB []Offer

type Help struct {
	OfferID       int64 `db:"offer_id"`
	ElaborationID int64 `db:"elaboration_id"`
}

// variables to store LEAGUES
var Helper Help

// help
