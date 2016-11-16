package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO replace boilerplate code with real one
type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", handleMovies).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func handleMovies(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var movies = map[string]*Movie{
		"tt0076759": {Title: "Star Wars: A New Hope", Rating: "8.7", Year: "1977"},
		"tt0082971": {Title: "Indiana Jones: Raiders of the Lost Ark", Rating: "8.6", Year: "1981"},
	}

	outgoingJSON, err := json.Marshal(movies)

	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}