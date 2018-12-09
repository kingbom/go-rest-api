package controller

import (
	"encoding/json"
	"net/http"

	. "github.com/kingbom/go-rest-api/model"
	. "github.com/kingbom/go-rest-api/repository"
	u "github.com/kingbom/go-rest-api/utils"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var movieRepo = MovieRepo{}

// GetMovies of movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := movieRepo.FindAll()
	if err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.RespondWithJson(w, http.StatusOK, movies)
}

// GetMovieById of movie
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := movieRepo.FindByID(params["id"])
	if err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	u.RespondWithJson(w, http.StatusOK, movie)
}

// CreateMovie a new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := movieRepo.Save(movie); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.RespondWithJson(w, http.StatusCreated, movie)
}

// UpdateMovie existing movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := movieRepo.Update(movie); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteMovie an existing movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		u.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := movieRepo.Delete(movie); err != nil {
		u.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	u.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
