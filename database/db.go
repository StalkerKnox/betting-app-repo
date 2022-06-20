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

	_, insertError := db.NamedExec(`INSERT INTO offers (number, tv_channel, id, title, has_statistics, time) VALUES (:number, :tv_channel, :id, :title, :has_statistics, :time)`, models.Offers)
	if insertError != nil {
		log.Fatal(insertError)
	}

	fmt.Println("succes")
}
