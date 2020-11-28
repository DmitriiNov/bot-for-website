package main

import (
	"net/http"
)

func getPort() (string, error) {
	return "1234", nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
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
