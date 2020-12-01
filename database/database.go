package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func StartDB() {
	dbc, err := sql.Open("sqlite3", "botdb.db")
	if err != nil {
		panic(err)
	}
	db = dbc
}

func AddToDB(form Form) {

}
