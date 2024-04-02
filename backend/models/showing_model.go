package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Showing struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	MovieShowingRef MovieRef           `bson:"movie,omitempty" json:"movie,omitempty" binding:"required"`
	Hall            Hall               `bson:"hall,omitempty" json:"hall,omitempty" binding:"required"`
	StartTime       time.Time          `bson:"startTime" json:"startTime" binding:"required"`
	EndTime         time.Time          `bson:"endTime" json:"endTime" binding:"required"`
	Seats           [][]Seat           `bson:"seats,omitempty" json:"seats,omitempty"`
	AvailableSeats  int                `bson:"availableSeats" json:"availableSeats" binding:"gte=0"`
	BookedSeats     int                `bson:"bookedSeats" json:"bookedSeats" binding:"gte=0"`
	PricePerTicket  float64            `bson:"pricePerTicket" json:"pricePerTicket" binding:"gte=0"`
	AudioType       string             `bson:"audioType" json:"audioType" binding:"required" validate:"required,eq=Dubbing|eq=Subtitles|eq=VoiceOver"`
	VideoType       string             `bson:"videoType" json:"videoType" binding:"required" validate:"required,eq=2D|eq=3D"`
}

type Seat struct {
	RowNumber  int  `bson:"rowNumber,omitempty" json:"rowNumber,omitempty"`
	SeatNumber int  `bson:"seatNumber,omitempty" json:"seatNumber,omitempty"`
	IsReserved bool `bson:"isReserved,omitempty" json:"isReserved,omitempty"`
}
