package main

import (
	"bot-for-website/bot"
	"bot-for-website/database"
	"encoding/json"
	"net/http"
)

func getPort() (string, error) {
	return "1234", nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var form database.Form
	_ = json.NewDecoder(r.Body).Decode(&form)
	go bot.SendForm(form)
	go database.AddToDB(form)
	w.WriteHeader(200)
}

func StartServer() error {

	port, err := getPort()
	if err != nil {
		return err
	}
	http.HandleFunc("/request", handler)
	go http.ListenAndServe(":"+port, nil)
	return nil
}
