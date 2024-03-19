package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Genre struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

type Movie struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Adult            bool               `json:"adult" bson:"adult"`
	Genres           []Genre            `json:"genres" bson:"genres"`
	Image            string             `json:"image" bson:"image"`
	Length           int                `json:"length" bson:"length"`
	OriginalLanguage string             `json:"originalLanguage" bson:"originalLanguage"`
	OriginalTitle    string             `json:"originalTitle" bson:"originalTitle"`
	Overview         string             `json:"overview" bson:"overview"`
	Popularity       float64            `json:"popularity" bson:"popularity"`
	ReleaseDate      time.Time          `json:"releaseDate" bson:"releaseDate"`
	Title            string             `json:"title" bson:"title"`
	VoteAverage      float64            `json:"voteAverage" bson:"voteAverage"`
	VoteCount        int                `json:"voteCount" bson:"voteCount"`
}

type MovieRef struct {
	MovieID    primitive.ObjectID `json:"movieId" bson:"movieId"`
	Categories []Genre            `json:"categories" bson:"categories"`
	Title      string             `json:"title" bson:"title"`
	Image      string             `json:"image" bson:"image"`
	Length     int                `json:"length" bson:"length"`
}

type Hall struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SeatsNumber int                `json:"seatsNumber" bson:"seatsNumber"`
	HallNumber  int                `json:"hallNumber" bson:"hallNumber"`
}

type Showing struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie     MovieRef           `json:"movie" bson:"movie"`
	Hall      Hall               `json:"hall" bson:"hall"`
	Datetime  time.Time          `json:"datetime" bson:"datetime"`
	FreeSeats int                `json:"freeSeats" bson:"freeSeats"`
	Type      string             `json:"type" bson:"type"`
	Price     float64            `json:"price" bson:"price"`
}

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type         string             `json:"type" bson:"type"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Firstname    string             `json:"firstname" bson:"firstname"`
	Lastname     string             `json:"lastname" bson:"lastname"`
	ToWatch      []MovieRef         `json:"to_watch" bson:"to_watch"`
	Reservations []Reservation      `json:"reservations" bson:"reservations"`
}

type Reservation struct {
	Showing Showing `json:"showing" bson:"showing"`
	Seat    int     `json:"seat" bson:"seat"`
}
