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

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updateEvent Event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "Please enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &updateEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updateEvent.Title
			singleEvent.Desc = updateEvent.Desc
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v was successfully deleted", eventID)
		}
	}
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Home")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := ":8080"
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/1event", getOneEvent)
	router.HandleFunc("/allevent", getAllEvents)
	fmt.Println("Listening server at port :8080")
	log.Fatal(http.ListenAndServe(port, router))
}
