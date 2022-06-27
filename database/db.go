package database

import (
	"betting-app/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:OvjAcbmOh4E@(localhost:3306)/betting_app")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var DB = ConnectDB()

func InsertToDB() error {

	for _, singleOffer := range models.Offers {
		_, insertErrOffers := DB.NamedExec(`INSERT INTO offers (number, tv_channel, offer_id, title, has_statistics, time) VALUES (:number, :tv_channel, :offer_id, :title, :has_statistics, :time)`, singleOffer)
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

func GetOffersFromDB() error {

	rows, _ := DB.Queryx("SELECT number, tv_channel, offer_id, title, has_statistics, time FROM offers")
	for rows.Next() {
		err := rows.StructScan(&models.SingleOffer)
		models.GetOffersFromDB = append(models.GetOffersFromDB, models.SingleOffer)
		if err != nil {
			return err
		}

	}

	for i, singleoffer := range models.GetOffersFromDB {

		rows, _ := DB.Queryx("SELECT offer_id, rate, name FROM rates WHERE offer_id = ? ", singleoffer.ID)
		for rows.Next() {

			err := rows.StructScan(&models.SingleRate)
			models.GetOffersFromDB[i].Rates = append(singleoffer.Rates, models.SingleRate)

			if err != nil {
				return err
			}
		}
		// models.GetOffersFromDB = append(models.GetOffersFromDB, singleoffer)

	}

	return nil
}

func GetOfferFromDB(req int) error {

	rows, _ := DB.Queryx("SELECT offer_id, rate, name FROM rates WHERE offer_id = ? ", req)
	for rows.Next() {
		err := rows.StructScan(&models.RateFromDB)
		models.OfferFromDB.Rates = append(models.OfferFromDB.Rates, models.RateFromDB)
		if err != nil {
			log.Fatalln(err)
		}
	}
	err := DB.Get(&models.OfferFromDB, "SELECT * FROM offers WHERE offer_id = ?", req)
	if err != nil {
		return err
	}
	return nil
}

func InsertOfferIntoDB(req models.Offer) error {
	_, insertErr := DB.NamedExec(`INSERT INTO offers (number, tv_channel, offer_id, title, has_statistics, time) VALUES (:number, :tv_channel, :offer_id, :title, :has_statistics, :time)`, req)
	if insertErr != nil {
		log.Fatal(insertErr)
	}

	for _, singleRate := range req.Rates {
		singleRate.OfferID = req.ID
		_, insertErr = DB.NamedExec(`INSERT INTO rates (offer_id, name, rate) VALUES (:offer_id, :name, :rate)`, singleRate)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
	}
	return nil
}

func InsertLeaguesIntoDB() {
	for _, singleLeague := range models.LeaguesStruct.Leagues {

		res, insertErr := DB.NamedExec(`INSERT INTO leagues (title) VALUES (:title)`, singleLeague)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
		leagueID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		for _, singleElaboration := range singleLeague.Elaborations {
			singleElaboration.LeagueID = leagueID
			res, insertErr := DB.NamedExec(`INSERT INTO elaborations (elaboration_id, league_id) VALUES (:elaboration_id, :league_id)`, singleElaboration)
			if insertErr != nil {
				log.Fatal(insertErr)
			}

			elaborationID, err := res.LastInsertId()
			if err != nil {
				log.Fatal(err)
			}

			for _, singleType := range singleElaboration.Types {
				singleType.ElaborationID = elaborationID
				_, insertErr := DB.NamedExec(`INSERT INTO types (elaboration_id, name) VALUES (:elaboration_id, :name)`, singleType)
				if insertErr != nil {
					log.Fatal(insertErr)
				}
			}
			for _, singleOffer := range singleElaboration.Offers {
				models.Helper.OfferID = singleOffer
				models.Helper.ElaborationID = elaborationID
				_, insertErr := DB.NamedExec(`INSERT INTO connect (elaboration_id, offer_id) VALUES (:elaboration_id, :offer_id)`, models.Helper)
				if insertErr != nil {
					log.Fatal(insertErr)
				}
			}

		}
	}
}

func GetLeaguesFromDB() error {
	rows, _ := DB.Queryx("SELECT title, id FROM leagues")
	for rows.Next() {
		err := rows.StructScan(&models.LeagueDB)
		if err != nil {
			return err
		}
		models.GetLeaguesFromDB.Leagues = append(models.GetLeaguesFromDB.Leagues, models.LeagueDB)
	}

	for i, singleLeague := range models.GetLeaguesFromDB.Leagues {
		rows, _ := DB.Queryx("SELECT elaboration_id FROM elaborations WHERE league_id = ?", singleLeague.ID)
		for rows.Next() {
			err := rows.StructScan(&models.ElaborationDB)
			if err != nil {
				return err
			}
			models.GetLeaguesFromDB.Leagues[i].Elaborations = append(models.GetLeaguesFromDB.Leagues[i].Elaborations, models.ElaborationDB)
		}

		for j, singleElaboration := range models.GetLeaguesFromDB.Leagues[i].Elaborations {
			rows, _ := DB.Queryx("SELECT name FROM types WHERE elaboration_id = ? ", singleElaboration.ID)
			for rows.Next() {
				err := rows.StructScan(&models.TypeDB)
				if err != nil {
					return err
				}
				models.GetLeaguesFromDB.Leagues[i].Elaborations[j].Types = append(models.GetLeaguesFromDB.Leagues[i].Elaborations[j].Types, models.TypeDB)
			}

			rows, _ = DB.Queryx("SELECT offer_id FROM connect WHERE elaboration_id = ? ", singleElaboration.ID)
			for rows.Next() {
				err := rows.StructScan(&models.OfferStruct)
				if err != nil {
					return err
				}
				models.GetLeaguesFromDB.Leagues[i].Elaborations[j].Offers = append(models.GetLeaguesFromDB.Leagues[i].Elaborations[j].Offers, models.OfferStruct.OfferID)
			}

		}

	}
	return nil

}
