package main

import (
	"log"
	"net/http"

	. "github.com/kingbom/go-rest-api/config"
	c "github.com/kingbom/go-rest-api/controller"
	. "github.com/kingbom/go-rest-api/repository"

	"github.com/gorilla/mux"
)

var config = Config{}
var movieRepo = MovieRepo{}

func init() {
	config.Read()

	movieRepo.Server = config.Server
	movieRepo.Database = config.Database
	movieRepo.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", c.GetMovies).Methods("GET")
	r.HandleFunc("/api/movies", c.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies", c.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies", c.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/movies/{id}", c.GetMovieById).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
