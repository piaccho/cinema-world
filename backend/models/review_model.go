package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId" binding:"required"`
	Firstname string             `json:"firstname" bson:"firstname" binding:"required"`
	Content   string             `json:"content" bson:"content" binding:"required,min=2,max=1000"`
	Rating    int                `json:"rating" bson:"rating" binding:"required,min=1,max=10"`
}
