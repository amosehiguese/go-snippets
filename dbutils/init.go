package dbutils

import (
	"database/sql"
	"log"
)

func Init(dbDriver *sql.DB) {
	statement, err := dbDriver.Prepare(station)

	if err != nil {
		log.Println(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Println("Table already exists")
	}
	log.Println("Created successfully!")
}