package database

import (
	"betting-app/models"
	"flag"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// CONNECT TO MYSQL DB
func ConnectDB() *sqlx.DB {
	connection := flag.String("connection", "root:OvjAcbmOh4E@(localhost:3307)/betting_app?parseTime=true", "default connection")
	flag.Parse()
	db, err := sqlx.Connect("mysql", *connection)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

// DEFINING GLOBAL VARIABLE DB
var DB = ConnectDB()

func InsertOffersIntoDB(x models.GetOfferResponse) error {

	for _, singleOffer := range x {
		var checker models.Offer
		_ = DB.Get(&checker, "SELECT * FROM offers WHERE offer_id = ?", singleOffer.ID)
		if checker.ID == singleOffer.ID {
			continue
		}

		_, insertErrOffers := DB.NamedExec(`INSERT offers (number, tv_channel, offer_id, title, has_statistics, time) VALUES (:number, :tv_channel, :offer_id, :title, :has_statistics, :time)`, singleOffer)
		if insertErrOffers != nil {
			return insertErrOffers
		}

		for _, singleRate := range singleOffer.Rates {
			singleRate.OfferID = singleOffer.ID
			_, insertErrRates := DB.NamedExec(`INSERT INTO rates (offer_id, name, rate) VALUES (:offer_id, :name, :rate)`, singleRate)
			if insertErrRates != nil {
				return insertErrRates
			}
		}

	}
	return nil
}

// INSERT LEAGUES INTO DB
func InsertLeaguesIntoDB(x models.GetLeagueResponse) error {
	var helper models.Help
	for _, singleLeague := range x.Leagues {
		var checker models.League
		_ = DB.Get(&checker, "SELECT * FROM leagues WHERE title = ?", singleLeague.Title)
		if checker.Title == singleLeague.Title {
			continue
		}

		res, insertErr := DB.NamedExec(`INSERT INTO leagues (title) VALUES (:title)`, singleLeague)
		if insertErr != nil {
			return insertErr
		}
		leagueID, err := res.LastInsertId()
		if err != nil {
			return err
		}

		for _, singleElaboration := range singleLeague.Elaborations {
			singleElaboration.LeagueID = leagueID
			res, insertErr := DB.NamedExec(`INSERT INTO elaborations (elaboration_id, league_id) VALUES (:elaboration_id, :league_id)`, singleElaboration)
			if insertErr != nil {
				return insertErr
			}

			elaborationID, err := res.LastInsertId()
			if err != nil {
				return err
			}

			for _, singleType := range singleElaboration.Types {
				singleType.ElaborationID = elaborationID
				_, insertErr := DB.NamedExec(`INSERT INTO types (elaboration_id, name) VALUES (:elaboration_id, :name)`, singleType)
				if insertErr != nil {
					return insertErr
				}
			}
			for _, singleOffer := range singleElaboration.Offers {
				helper.OfferID = singleOffer
				helper.ElaborationID = elaborationID
				_, insertErr := DB.NamedExec(`INSERT INTO connect (elaboration_id, offer_id) VALUES (:elaboration_id, :offer_id)`, helper)
				if insertErr != nil {
					return insertErr
				}
			}

		}
	}
	return nil
}

// GET OFFERS FROM DB
func GetOffersFromDB() (*[]models.Offer, error) {
	var getOffersFromDB []models.Offer
	var oneOffer models.Offer
	var singleRate models.Rate
	rows, _ := DB.Queryx("SELECT number, tv_channel, offer_id, title, has_statistics, time FROM offers")
	for rows.Next() {
		err := rows.StructScan(&oneOffer)
		getOffersFromDB = append(getOffersFromDB, oneOffer)
		if err != nil {
			return nil, err
		}

	}

	for i, singleOffer := range getOffersFromDB {

		rows, _ := DB.Queryx("SELECT offer_id, rate, name FROM rates WHERE offer_id = ? ", singleOffer.ID)
		for rows.Next() {

			err := rows.StructScan(&singleRate)
			getOffersFromDB[i].Rates = append(getOffersFromDB[i].Rates, singleRate)

			if err != nil {
				return nil, err
			}
		}

	}

	return &getOffersFromDB, nil
}

//GET SINGLE OFFER FROM DB
func GetOfferFromDB(req int) (*models.Offer, error) {
	var offerFromDB models.Offer
	var rateFromDB models.Rate
	rows, _ := DB.Queryx("SELECT offer_id, rate, name FROM rates WHERE offer_id = ? ", req)
	for rows.Next() {
		err := rows.StructScan(&rateFromDB)
		offerFromDB.Rates = append(offerFromDB.Rates, rateFromDB)
		if err != nil {
			return nil, err
		}
	}
	err := DB.Get(&offerFromDB, "SELECT * FROM offers WHERE offer_id = ?", req)
	if err != nil {
		return nil, err
	}
	return &offerFromDB, nil
}

// INSERT SINGLE OFFER INTO DB
func InsertOfferIntoDB(req models.Offer) error {
	_, insertErr := DB.NamedExec(`INSERT INTO offers (number, tv_channel, offer_id, title, has_statistics, time) VALUES (:number, :tv_channel, :offer_id, :title, :has_statistics, :time)`, req)
	if insertErr != nil {
		return insertErr
	}

	for _, singleRate := range req.Rates {
		singleRate.OfferID = req.ID
		_, insertErr = DB.NamedExec(`INSERT INTO rates (offer_id, name, rate) VALUES (:offer_id, :name, :rate)`, singleRate)
		if insertErr != nil {
			return insertErr
		}
	}
	return nil
}

// GET LEAGUES FROM DB
func GetLeaguesFromDB() (*models.GetLeagueResponse, error) {
	var getLeaguesFromDB models.GetLeagueResponse
	var leagueDB models.League
	rows, _ := DB.Queryx("SELECT title, id FROM leagues")
	for rows.Next() {
		err := rows.StructScan(&leagueDB)
		if err != nil {
			return nil, err
		}
		getLeaguesFromDB.Leagues = append(getLeaguesFromDB.Leagues, leagueDB)
	}

	var elaborationDB models.Elaboration
	for i, singleLeague := range getLeaguesFromDB.Leagues {
		rows, _ := DB.Queryx("SELECT elaboration_id FROM elaborations WHERE league_id = ?", singleLeague.ID)
		for rows.Next() {
			err := rows.StructScan(&elaborationDB)
			if err != nil {
				return nil, err
			}
			getLeaguesFromDB.Leagues[i].Elaborations = append(getLeaguesFromDB.Leagues[i].Elaborations, elaborationDB)
		}

		var typeDB models.Type
		for j, singleElaboration := range getLeaguesFromDB.Leagues[i].Elaborations {
			rows, _ := DB.Queryx("SELECT name FROM types WHERE elaboration_id = ? ", singleElaboration.ID)
			for rows.Next() {
				err := rows.StructScan(&typeDB)
				if err != nil {
					return nil, err
				}
				getLeaguesFromDB.Leagues[i].Elaborations[j].Types = append(getLeaguesFromDB.Leagues[i].Elaborations[j].Types, typeDB)
			}

			var offerStruct models.Help
			rows, _ = DB.Queryx("SELECT offer_id FROM connect WHERE elaboration_id = ? ", singleElaboration.ID)
			for rows.Next() {
				err := rows.StructScan(&offerStruct)
				if err != nil {
					return nil, err
				}
				getLeaguesFromDB.Leagues[i].Elaborations[j].Offers = append(getLeaguesFromDB.Leagues[i].Elaborations[j].Offers, offerStruct.OfferID)
			}

		}

	}
	return &getLeaguesFromDB, nil

}

// GET USER BALANCE FROM DB
func GetBalanceFromDB(x models.TikcetDesign) (float64, error) {
	err := DB.Get(&x.RemainingBalance, "SELECT balance FROM players WHERE user_name = ?", x.UserName)
	if err != nil {
		return 0, err
	}
	return x.RemainingBalance, nil

}

// GET RATES FOR EVERY SINGLE PLAYED TYPE
func GetRatesFromDB(x models.TikcetDesign) (*models.TikcetDesign, error) {
	var calculatorStorage []float64
	var playOffer models.PlayedOffer

	for i, singlePlayedOffer := range x.PlayedOffers {
		rows, _ := DB.Queryx("SELECT * FROM rates WHERE offer_id = ? AND name = ?", singlePlayedOffer.ID, singlePlayedOffer.Name)
		for rows.Next() {
			err := rows.StructScan(&playOffer)
			if err != nil {
				return nil, err
			}
		}
		x.PlayedOffers[i].Rate = playOffer.Rate
		calculatorStorage = append(calculatorStorage, playOffer.Rate)

		var prizeMoney float64 = x.PaymentAmount
		for _, coef := range calculatorStorage {
			prizeMoney = coef * prizeMoney
		}
		x.PrizeMoney = prizeMoney

	}
	x.RemainingBalance = x.RemainingBalance - x.PaymentAmount

	return &x, nil
}

// UPDATE user balance after simulating ticket paying
func UpdateBalance(x models.TikcetDesign) error {
	_, updateErr := DB.Exec(`UPDATE players SET balance = ? WHERE user_name = ?`, x.RemainingBalance, x.UserName)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func InsertTicketIntoDB(x models.TikcetDesign) (*models.TikcetDesign, error) {
	res, insertErr := DB.NamedExec(`INSERT INTO tickets (ticket_id, user_name, payment_amount, prize_money) VALUES (:ticket_id, :user_name, :payment_amount, :prize_money)`, x)
	if insertErr != nil {
		return nil, insertErr
	}

	ticketID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	x.ID = int(ticketID)

	for _, singleOffer := range x.PlayedOffers {
		singleOffer.TicketID = x.ID
		_, insertErr = DB.NamedExec(`INSERT INTO played_offers (ticket_id, offer_id, rate, name) VALUES (:ticket_id, :offer_id, :rate, :name)`, singleOffer)
		if insertErr != nil {
			return nil, insertErr
		}
	}
	return &x, nil
}
