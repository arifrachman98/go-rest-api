package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}

type allEvents []Event

var events = allEvents{
	{
		ID:    "1",
		Title: "Golang Tesing REST api",
		Desc:  "In this chapter, you will learn about something good",
	},
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "Please enter data with the event title and description olny in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Home")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := ":8080"
	router.HandleFunc("/", homeLink)
	fmt.Println("Listening server at port :8080")
	log.Fatal(http.ListenAndServe(port, router))
}
