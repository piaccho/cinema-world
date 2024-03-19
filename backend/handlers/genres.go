package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetGenres(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := client.Database("cinema").Collection("genres")

		// Find all documents in the collection
		cursor, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Decode the documents
		var genres []bson.M
		if err = cursor.All(context.Background(), &genres); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, genres)
	}
}

func CreateGenre(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := client.Database("cinema").Collection("genres")

		// Create a new document
		var genre bson.M
		if err := c.BindJSON(&genre); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert the document
		if _, err := collection.InsertOne(context.Background(), genre); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the document as JSON
		c.JSON(http.StatusCreated, genre)
	}
}
