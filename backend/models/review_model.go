package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId" binding:"required"`
	FirstName string             `json:"firstName" bson:"firstName" binding:"required"`
	Content   string             `json:"content" bson:"content" binding:"required"`
}
