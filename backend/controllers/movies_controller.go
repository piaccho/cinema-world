package controllers

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"piaccho/cinema-api/configs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"piaccho/cinema-api/models"
)

var moviesCollection *mongo.Collection = configs.GetCollection("movies")

// DefaultQuantity Default number of documents to return
const DefaultQuantity int64 = int64(20)

// GetMovies returns a list of movies based on the type - popular, upcoming, genres, search
func GetMovies(getType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter bson.M
		findOptions := options.Find()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var movies []models.Movie
		defer cancel()

		switch getType {
		case "popular":
			filter = bson.M{}
			// Get quantity and convert to an int64
			quantity, err := strconv.ParseInt(c.Param("quantity"), 10, 64)
			if err != nil {
				quantity = DefaultQuantity
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
				quantity = DefaultQuantity
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
		results, err := moviesCollection.Find(ctx, filter, findOptions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// reading from the db in an optimal way
		defer func(results *mongo.Cursor, ctx context.Context) {
			err := results.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
		}(results, ctx)

		for results.Next(ctx) {
			var singleMovie models.Movie
			if err = results.Decode(&singleMovie); err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			movies = append(movies, singleMovie)
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": movies}},
		)
	}
}

func CreateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var movie models.Movie
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Insert a single document
		result, err := moviesCollection.InsertOne(ctx, movie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result.InsertedID}},
		)
	}
}

func GetMovieById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var movie models.Movie
		defer cancel()

		// Get movie id
		movieId, _ := primitive.ObjectIDFromHex(c.Param("movieId"))

		// Find a single document
		err := moviesCollection.FindOne(ctx, bson.M{"_id": movieId}).Decode(&movie)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No movie found with the provided ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": movie}},
		)
	}
}

func UpdateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var movie models.Movie
		defer cancel()

		// Get movie id
		movieId, _ := primitive.ObjectIDFromHex(c.Param("movieId"))

		// validate the request body
		if err := c.BindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// Update a single document
		result, err := moviesCollection.UpdateOne(ctx, bson.M{"_id": movieId}, bson.M{"$set": movie})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// get updated movie details
		var updatedMovie models.Movie
		if result.MatchedCount == 1 {
			err := moviesCollection.FindOne(ctx, bson.M{"_id": movieId}).Decode(&updatedMovie)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedMovie}},
		)
	}
}

func DeleteMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get movie id
		movieId, _ := primitive.ObjectIDFromHex(c.Param("movieId"))

		// Delete a single document
		result, err := moviesCollection.DeleteOne(ctx, bson.M{"_id": movieId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Movie with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Movie successfully deleted!"}},
		)
	}
}

func GetMoviesRefs() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var movies []models.MovieRef
		defer cancel()

		// Find all documents in the collection
		results, err := moviesCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// reading from the db in an optimal way
		defer func(results *mongo.Cursor, ctx context.Context) {
			err := results.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
		}(results, ctx)

		for results.Next(ctx) {
			var singleMovie models.Movie
			if err = results.Decode(&singleMovie); err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			movies = append(movies, models.MovieRef{
				Id:     singleMovie.Id,
				Genres: singleMovie.Genres,
				Title:  singleMovie.Title,
				Image:  singleMovie.Image,
				Length: singleMovie.Length,
			})
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": movies}},
		)
	}
}
