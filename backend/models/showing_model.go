package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Showing struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	MovieShowingRef MovieRef           `bson:"movie,omitempty" json:"movie,omitempty" binding:"required"`
	Hall            Hall               `bson:"hall,omitempty" json:"hall,omitempty" binding:"required"`
	StartTime       time.Time          `bson:"start_time" json:"start_time" binding:"required"`
	EndTime         time.Time          `bson:"end_time" json:"end_time" binding:"required"`
	Seats           [][]Seat           `bson:"seats,omitempty" json:"seats,omitempty"`
	AvailableSeats  int                `bson:"available_seats" json:"available_seats" binding:"gte=0"`
	BookedSeats     int                `bson:"booked_seats" json:"booked_seats" binding:"gte=0"`
	PricePerTicket  float64            `bson:"price_per_ticket" json:"price_per_ticket" binding:"gte=0"`
	AudioType       string             `bson:"audio_type" json:"audio_type" binding:"required" validate:"required,eq=Dubbing|eq=Subtitles|eq=VoiceOver"`
	VideoType       string             `bson:"video_type" json:"video_type" binding:"required" validate:"required,eq=2D|eq=3D"`
}

type Seat struct {
	RowNumber  int  `bson:"row_number,omitempty" json:"row_number,omitempty"`
	SeatNumber int  `bson:"seat_number,omitempty" json:"seat_number,omitempty"`
	IsReserved bool `bson:"is_reserved,omitempty" json:"is_reserved,omitempty"`
}
