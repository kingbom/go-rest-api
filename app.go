package main

import (
	"log"
	"net/http"

	. "github.com/kingbom/go-rest-api/config"
	"github.com/kingbom/go-rest-api/controller"
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
	r.HandleFunc("/api/movies", controller.GetMovies).Methods("GET")
	r.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies", controller.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/movies/{id}", controller.GetMovieById).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
