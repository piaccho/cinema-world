package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ShowingId   primitive.ObjectID `bson:"showing_id,omitempty" json:"showing_id,omitempty"`
	UserId      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	SeatNumbers []string           `bson:"seat_numbers" json:"seat_numbers" binding:"required"`
	TotalPrice  float64            `bson:"total_price" json:"total_price" binding:"gte=0"`
}
