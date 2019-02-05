package main

import (
	"github.com/gorilla/mux"
	"github/pedroportasvieira/go-api/tags"
	"log"
	"net/http"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tag", tags.Tags).Methods("GET")
	router.HandleFunc("/tag", tags.SaveTag).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}