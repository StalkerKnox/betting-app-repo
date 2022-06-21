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
		res, insertError := db.NamedExec(`INSERT INTO offers (number, tv_channel, id, title, has_statistics, time) VALUES (:number, :tv_channel, :id, :title, :has_statistics, :time)`, offer)
		if insertError != nil {
			log.Fatal(insertError)
		}

		lastId, err1 := res.LastInsertId()
		if err1 != nil {
			log.Fatal(err1)
		}

		for _, rate := range offer.Rates {
			rate.Id_rate = int(lastId)
			_, inserterr1 := db.NamedExec(`INSERT INTO rates (id_rate, name, rate) VALUES (:id_rate, :name, :rate)`, rate)
			if inserterr1 != nil {
				log.Fatal(inserterr1)
			}
		}
	}

	fmt.Println("succes")
}
