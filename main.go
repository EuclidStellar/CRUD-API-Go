package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {

	ID string `json:"id"`
	Title string `json:"title"`
	Year string `json:"year"`
	Director Director `json:"director"`


}

type Director struct {

	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Age int `json:"age"`


}


var movies []Movie

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}