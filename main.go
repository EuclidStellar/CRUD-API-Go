package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {

	ID string `json:"id"`
	Title string `json:"title"`
	Year string `json:"year"`
	Director *Director `json:"director"`


}

type Director struct {

	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Age int `json:"age"`


}


var movies []Movie


func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return

		}

	}

	json.NewEncoder(w).Encode(&Movie{})

}




func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "The Shawshank Redemption", Year: "1994", Director: &Director{FirstName: "Frank", LastName: "Darabont", Age: 61}})
	movies = append(movies, Movie{ID: "2", Title: "The Godfather", Year: "1972", Director: &Director{FirstName: "Francis", LastName: "Ford Coppola", Age: 81}})
	movies = append(movies , Movie{ID: "3", Title: "The Dark Knight", Year: "2008", Director: &Director{FirstName: "Christopher", LastName: "Nolan", Age: 50}})
	movies = append(movies, Movie{ID: "4" , Title : "Swadesh" , Year : "2004" , Director : &Director{FirstName : "Ashutosh" , LastName : "Gowariker" , Age : 56}})
	movies = append(movies , Movie{ID : "5" , Title : "The Pursuit of Happyness" , Year : "2006" , Director : &Director{FirstName : "Gabriele" , LastName : "Muccino" , Age : 53}})


	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}