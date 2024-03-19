package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"piaccho/cinema-api/models"
)

func GetMovies(client *mongo.Client, getType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := client.Database("cinema").Collection("movies")
		var filter bson.M
		findOptions := options.Find()
		// Default number of documents to return
		DEFAULT_QUANTITY := int64(20)

		switch getType {
		case "popular":
			filter = bson.M{}
			// Get quantity and convert to an int64
			quantity, err := strconv.ParseInt(c.Param("quantity"), 10, 64)
			if err != nil {
				quantity = DEFAULT_QUANTITY
				// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
				// return
			}
			findOptions.SetSort(bson.D{{"voteCount", -1}})
			findOptions.SetLimit(quantity)

		case "upcoming":
			filter = bson.M{}
			// Get quantity and convert to an int64
			quantity, err := strconv.ParseInt(c.Param("quantity"), 10, 64)
			if err != nil {
				quantity = DEFAULT_QUANTITY
				// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
				// return
			}
			findOptions.SetSort(bson.D{{"releaseDate", -1}})
			findOptions.SetLimit(quantity)

			// releaseDate greater than or equal to today
			// filter = bson.M{"releaseDate": bson.M{"$gte": time.Now().Format(time.RFC3339)}}

		case "genres":
			// Get genre name
			genreName := c.Param("name")

			filter = bson.M{"genres.name": genreName}

		case "search":
			query := c.Param("query")

			filter = bson.M{"title": bson.M{"$regex": query, "$options": "i"}}

		default:
			filter = bson.M{}
		}

		// Find all documents in the collection
		cursor, err := collection.Find(context.Background(), filter, findOptions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Decode the documents
		var movies []models.Movie
		if err = cursor.All(context.Background(), &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the documents as JSON
		c.JSON(http.StatusOK, movies)
	}
}
