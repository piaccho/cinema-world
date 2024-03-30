package controllers

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"piaccho/cinema-api/configs"
	"piaccho/cinema-api/models"
	"piaccho/cinema-api/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var showingCollection = configs.GetCollection("showings")

func GetAllShowings() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var showings []models.Showing
		defer cancel()

		results, err := showingCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// reading from the db in an optimal way
		defer func(results *mongo.Cursor, ctx context.Context) {
			err := results.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
		}(results, ctx)

		for results.Next(ctx) {
			var singleShowing models.Showing
			if err = results.Decode(&singleShowing); err != nil {
				c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			showings = append(showings, singleShowing)
		}

		c.JSON(http.StatusOK,
			models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": showings}},
		)
	}
}

func CreateShowing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var showing models.Showing
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&showing); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&showing); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newShowing := models.Showing{
			Id:             primitive.NewObjectID(),
			MovieID:        showing.MovieID,
			HallId:         showing.HallId,
			StartTime:      showing.StartTime,
			EndTime:        showing.EndTime,
			AvailableSeats: showing.AvailableSeats,
			BookedSeats:    showing.BookedSeats,
			PricePerTicket: showing.PricePerTicket,
			AudioType:      showing.AudioType,
			VideoType:      showing.VideoType,
		}

		result, err := showingCollection.InsertOne(ctx, newShowing)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, models.ResponseModel{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetShowingById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		showingId := c.Param("showingId")
		var showing models.Showing
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(showingId)

		err := showingCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&showing)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.ResponseModel{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No showing found with the provided ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": showing}})
	}
}

func UpdateShowing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		showingId := c.Param("showingId")
		var showing models.Showing
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(showingId)

		// validate the request body
		if err := c.BindJSON(&showing); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&showing); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"movie_id": showing.MovieID, "hall_id": showing.HallId, "start_time": showing.StartTime, "end_time": showing.EndTime, "available_seats": showing.AvailableSeats, "booked_seats": showing.BookedSeats, "price_per_ticket": showing.PricePerTicket, "audio_type": showing.AudioType, "video_type": showing.VideoType}
		result, err := showingCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// get updated showing details
		var updatedShowing models.Showing
		if result.MatchedCount == 1 {
			err := showingCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedShowing)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedShowing}})
	}
}

func DeleteShowing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		showingId := c.Param("showingId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(showingId)

		result, err := showingCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				models.ResponseModel{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Showing with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Showing successfully deleted!"}},
		)
	}
}

func GetShowingsByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse datetime from the URL parameter
		datetimeStr := c.Param("datetime")
		datetime, err := time.Parse(time.RFC3339, datetimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format. Please use RFC3339."})
			return
		}

		// Find all documents in the showingCollection that have the given datetime
		cursor, err := showingCollection.Find(context.Background(), bson.D{{"datetime", datetime}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Decode the documents
		var showings []models.Showing
		if err = cursor.All(context.Background(), &showings); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, showings)
	}
}

func GetShowingsByMovieId() gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := c.Param("movieId")

		// Find all documents in the showingCollection that contain the movie with the given id
		cursor, err := showingCollection.Find(context.Background(), bson.D{{"movie._id", movieId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Decode the documents
		var showings []models.Showing
		if err = cursor.All(context.Background(), &showings); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, showings)
	}
}

func GetShowingsByHallId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hallId := c.Param("hallId")
		defer cancel()

		// Find all documents in the showingCollection that contain the hall with the given id
		cursor, err := showingCollection.Find(ctx, bson.D{{"hall._id", hallId}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Decode the documents
		var showings []models.Showing
		if err = cursor.All(ctx, &showings); err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": showings}})
	}
}

var audioTypes = []string{"Dubbing", "Subtitles", "VoiceOver"}
var videoTypes = []string{"2D", "3D"}

//func GenerateShowingsForNextDays() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//		daysNumber := c.Param("daysNumber")
//		defer cancel()
//
//		// Parse the daysNumber from the URL parameter
//		days, err := primitive.ParseDecimal128(daysNumber)
//		if err != nil {
//			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid days number format. Please use a number."}})
//			return
//		}
//
//		// Get the current date
//		currentDate := time.Now()
//
//		// Generate showings for the next days
//		for i := 0; i < days; i++ {
//			// Create a new showing
//			newShowing := models.Showing{
//				Id:             primitive.NewObjectID(),
//				MovieID:        primitive.NewObjectID(),
//				HallId:         primitive.NewObjectID(),
//				StartTime:      currentDate.AddDate(0, 0, i),
//				EndTime:        currentDate.AddDate(0, 0, i).Add(time.Hour * 2),
//				AvailableSeats: 100,
//				BookedSeats:    0,
//				PricePerTicket: 10.0,
//				AudioType:      "Dubbing",
//				VideoType:      "2D",
//			}
//
//			// Insert the new showing into the database
//			_, err := showingCollection.InsertOne(ctx, newShowing)
//			if err != nil {
//				c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
//				return
//			}
//		}
//
//		c.JSON(http.StatusCreated, models.ResponseModel{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": "Showings successfully generated for the next " + daysNumber + " days."}})
//	}
//}
