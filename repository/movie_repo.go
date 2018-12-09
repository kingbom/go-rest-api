package repository

import (
	"log"

	. "github.com/kingbom/go-rest-api/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MovieRepo struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

// Conect to database
func (m *MovieRepo) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//FindAll of movies
func (m *MovieRepo) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

//FindByID  movie
func (m *MovieRepo) FindByID(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.IsObjectIdHex(id)).One(&movie)
	return movie, err
}

//Save movie
func (m *MovieRepo) Save(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

//Update movie
func (m *MovieRepo) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}

//Delete movie
func (m *MovieRepo) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}
