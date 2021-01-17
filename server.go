package main

import (
	"bot-for-website/bot"
	"bot-for-website/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPort() (string, error) {
	return "8080", nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var form database.Form
	bodyb, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err := json.Unmarshal(bodyb, &form)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(200)
		return
	}
	go bot.SendForm(form)
	go database.AddToDB(form)
	w.WriteHeader(200)
}

func StartServer() error {

	port, err := getPort()
	if err != nil {
		return err
	}
	database.StartDB()
	http.HandleFunc("/request", handler)
	go http.ListenAndServe(":"+port, nil)
	return nil
}
