package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetHalls(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := client.Database("cinema").Collection("halls")

		// Find all documents in the collection
		cursor, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Decode the documents
		var halls []bson.M
		if err = cursor.All(context.Background(), &halls); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, halls)
	}
}
