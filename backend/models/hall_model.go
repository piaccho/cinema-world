package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hall struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name" json:"name" binding:"required"`
	Rows        int                `bson:"rows" json:"rows" binding:"required,gte=1"`
	SeatsPerRow int                `bson:"seats_per_row" json:"seats_per_row" binding:"required,gte=1"`
}
