package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Showing struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	MovieShowingRef MovieRef           `bson:"movie_id,omitempty" json:"movie_id,omitempty" binding:"required"`
	HallId          primitive.ObjectID `bson:"hall_id,omitempty" json:"hall_id,omitempty" binding:"required"`
	StartTime       time.Time          `bson:"start_time" json:"start_time" binding:"required"`
	EndTime         time.Time          `bson:"end_time" json:"end_time" binding:"required"`
	AvailableSeats  int                `bson:"available_seats" json:"available_seats" binding:"gte=0"`
	BookedSeats     int                `bson:"booked_seats" json:"booked_seats" binding:"gte=0"`
	PricePerTicket  float64            `bson:"price_per_ticket" json:"price_per_ticket" binding:"gte=0"`
	AudioType       string             `bson:"audio_type" json:"audio_type" binding:"required" validate:"required,eq=Dubbing|eq=Subtitles|eq=VoiceOver"`
	VideoType       string             `bson:"video_type" json:"video_type" binding:"required" validate:"required,eq=2D|eq=3D"`
}
