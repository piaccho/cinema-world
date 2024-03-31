package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      string             `json:"type" bson:"type" validate:"required,eq=Admin|eq=User"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Password  string             `json:"password" bson:"password" validate:"required,min=8,max=64"`
	Firstname string             `json:"firstname" bson:"firstname" validate:"required,min=2"`
	Lastname  string             `json:"lastname" bson:"lastname" validate:"required,min=2"`
	ToWatch   []ToWatchListItem  `json:"to_watch" bson:"to_watch"`
}

type ToWatchListItem struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie MovieRef           `json:"movie" bson:"movie"`
}
