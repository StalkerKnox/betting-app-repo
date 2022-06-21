package database

import (
	"betting-app/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() {
	db, err := sqlx.Connect("mysql", "root:OvjAcbmOh4E@(localhost:3306)/betting_app")
	if err != nil {
		log.Fatal(err)
	}

	for _, offer := range models.Offers {
		_, insertError := db.NamedExec(`INSERT INTO offers (number, tv_channel, offer_id, title, has_statistics, time) VALUES (:number, :tv_channel, :offer_id, :title, :has_statistics, :time)`, offer)
		if insertError != nil {
			log.Fatal(insertError)
		}

		offerID := offer.ID

		for _, rate := range offer.Rates {
			rate.OfferID = offerID
			_, inserterr1 := db.NamedExec(`INSERT INTO rates (offer_id, name, rate) VALUES (:offer_id, :name, :rate)`, rate)
			if inserterr1 != nil {
				log.Fatal(inserterr1)
			}
		}
	}

	fmt.Println("succes")
}
