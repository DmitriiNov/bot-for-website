package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

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

func GetLast10() ([]Form, []int, error) {
	var id int
	var name, text, service, contacts string
	rows, err := db.Query("SELECT id, name, info, service, contacts FROM forms ORDER BY id DESC LIMIT 10")
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var data []Form
	var ids []int
	for rows.Next() {
		err := rows.Scan(&id, &name, &text, &service, &contacts)
		if err != nil {
			return nil, nil, err
		}
		var conts []Contact
		json.Unmarshal([]byte(contacts), &conts)
		data = append(data, Form{
			Name:     name,
			Info:     text,
			Service:  service,
			Contacts: conts,
		})
		ids = append(ids, id)

	}
	err = rows.Err()
	if err != nil {
		return nil, nil, err
	}

	return data, ids, nil
}

func AddToDB(form Form) {
	cons, err := json.Marshal(form.Contacts)
	if err != nil {
		return
	}
	stmt, err := db.Prepare("INSERT INTO forms (name, service, contacts, info) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(form.Name, form.Service, cons, form.Info)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Affect %d\n", lastId)
}
