package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Home")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := ":8080"
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(port, router))
}
