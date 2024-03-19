package handlers

import (
	"context"
	"net/http"
	"piaccho/cinema-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetShowingsByDate(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse datetime from the URL parameter
		datetimeStr := c.Param("datetime")
		datetime, err := time.Parse(time.RFC3339, datetimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format. Please use RFC3339."})
			return
		}

		collection := client.Database("cinema").Collection("showings")

		// Find all documents in the collection that have the given datetime
		cursor, err := collection.Find(context.Background(), bson.D{{"datetime", datetime}})
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

func GetShowingsByMovieId(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := c.Param("id")

		collection := client.Database("cinema").Collection("showings")

		// Find all documents in the collection that contain the movie with the given id
		cursor, err := collection.Find(context.Background(), bson.D{{"movie._id", movieId}})
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
