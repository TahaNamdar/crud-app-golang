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
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie


func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}


func main(){
	r:= mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn: "24342",Title: "Movie 1",Director: &Director{Firstname: "ali",Lastname: "namdar"}})
	movies = append(movies, Movie{ID:"2", Isbn: "4353",Title: "Movie 2",Director: &Director{Firstname: "taha",Lastname: "javan"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")

	fmt.Printf("starting on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

}