package database

import (
	"betting-app/models"
	"fmt"
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
		models.IterationStorage = append(models.IterationStorage, models.SingleOffer)
		if err != nil {
			return err
		}

	}

	for _, singleoffer := range models.IterationStorage {
		var rows *sqlx.Rows
		rows, _ = DB.Queryx("SELECT offer_id, rate, name FROM rates WHERE offer_id = ? ", singleoffer.ID)
		for rows.Next() {

			err := rows.StructScan(&models.SingleRate)
			singleoffer.Rates = append(singleoffer.Rates, models.SingleRate)

			if err != nil {
				return err
			}
		}
		models.GetOffersFromDB = append(models.GetOffersFromDB, singleoffer)

	}

	// fmt.Println(models.GetOffersFromDB)
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
		// rows, _ := DB.Queryx("SELECT id FROM leagues")
		// for rows.Next() {
		// 	err := rows.StructScan(&singleLeague)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// }

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
			// rows, _ := DB.Queryx("SELECT elaboration_id FROM elaborations")
			// for rows.Next() {
			// 	err := rows.StructScan(&singleElaboration)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// }
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

func GetLeaguesFromDB() {
	rows, _ := DB.Queryx("SELECT title, id FROM leagues")
	for rows.Next() {
		err := rows.StructScan(&models.LeagueDB)
		if err != nil {
			log.Fatal(err)
		}
		models.LeaguesStructDB.Leagues = append(models.LeaguesStructDB.Leagues, models.LeagueDB)
	}
	fmt.Println(models.LeaguesStructDB.Leagues)

	for _, singleLeague := range models.LeaguesStructDB.Leagues {
		rows, _ := DB.Queryx("SELECT elaboration_id FROM elaborations WHERE league_id = ?", singleLeague.ID)
		for rows.Next() {
			err := rows.StructScan(&models.ElaborationDB)
			if err != nil {
				log.Fatal(err)
			}
			singleLeague.Elaborations = append(singleLeague.Elaborations, models.ElaborationDB)
		}

		for _, singleElaboration := range singleLeague.Elaborations {
			rows, _ := DB.Queryx("SELECT name FROM types WHERE elaboration_id = ? ", singleElaboration.ID)
			for rows.Next() {
				err := rows.StructScan(&models.TypeDB)
				if err != nil {
					log.Fatal(err)
				}
				singleElaboration.Types = append(singleElaboration.Types, models.TypeDB)
			}

			fmt.Println(singleElaboration.Types)
		}
		models.LeaguesStructDBDB.Leagues = append(models.LeaguesStructDBDB.Leagues, singleLeague)
	}

	fmt.Println(models.LeaguesStructDBDB)

}
