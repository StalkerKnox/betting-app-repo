package models

import "time"

// Creating GetLeagueResponse Struct
type GetLeagueResponse struct {
	Leagues []League `json:"lige"`
}

type League struct {
	Title        string        `json:"naziv" db:"title"`
	Elaborations []Elaboration `json:"razrade"`
	ID           int           `json:"-" db:"id"`
}

type Elaboration struct {
	Types    []Type  `json:"tipovi" `
	Offers   []int64 `json:"ponude"`
	LeagueID int64   `json:"-" db:"league_id"`
	ID       int64   `json:"-" db:"elaboration_id"`
}

type Type struct {
	Name          string `json:"naziv" db:"name"`
	ElaborationID int64  `json:"-" db:"elaboration_id"`
	ID            int64  `json:"-" db:"type_id"`
}

//Creating GetOfferResponse Slice of Structs
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
	OfferID int     `json:"-" db:"offer_id"`
	Rate    float64 `json:"tecaj" db:"rate" validate:"required"`
	Name    string  `json:"naziv" db:"name" validate:"required"`
}

//DATABASE

// type OffersFromDB []Offer

// Ticket design schema
type TikcetDesign struct {
	ID               int           `db:"ticket_id"`
	UserName         string        `json:"korisnicko_ime" db:"user_name"`
	PaymentAmount    float64       `json:"uplaceni_iznos" db:"payment_amount"`
	PlayedOffers     []PlayedOffer `json:"lista_odigranih_ponuda"`
	PrizeMoney       float64       `json:"moguci_dobitak" db:"prize_money"`
	RemainingBalance float64       `json:"-" db:"balance"`
}

type PlayedOffer struct {
	ID       int     `json:"id_ponude" db:"offer_id"`
	TicketID int     `json:"-" db:"ticket_id"`
	Name     string  `json:"odigrani_tip" db:"name"`
	Rate     float64 `json:"-" db:"rate"`
}

// Helper struct for DB operations
type Help struct {
	OfferID       int64 `db:"offer_id"`
	ElaborationID int64 `db:"elaboration_id"`
}
