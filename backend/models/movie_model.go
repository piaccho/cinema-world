package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Movie struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Adult            bool               `json:"adult" bson:"adult" binding:"required"`
	Genres           []Genre            `json:"genres" bson:"genres" binding:"required"`
	Image            string             `json:"image" bson:"image" binding:"required"`
	Length           int                `json:"length" bson:"length" binding:"required"`
	OriginalLanguage string             `json:"originalLanguage" bson:"originalLanguage" binding:"required"`
	OriginalTitle    string             `json:"originalTitle" bson:"originalTitle" binding:"required"`
	Overview         string             `json:"overview" bson:"overview" binding:"required"`
	Popularity       float64            `json:"popularity" bson:"popularity"`
	ReleaseDate      time.Time          `json:"releaseDate" bson:"releaseDate" binding:"required"`
	Title            string             `json:"title" bson:"title" binding:"required"`
	VoteAverage      float64            `json:"voteAverage" bson:"voteAverage"`
	VoteCount        int                `json:"voteCount" bson:"voteCount"`
	Reviews          []Review           `json:"reviews" bson:"reviews"`
}

type MovieRef struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Genres []Genre            `json:"genres" bson:"genres" binding:"required"`
	Image  string             `json:"image" bson:"image" binding:"required"`
	Length int                `json:"length" bson:"length" binding:"required"`
	Title  string             `json:"title" bson:"title" binding:"required"`
}
