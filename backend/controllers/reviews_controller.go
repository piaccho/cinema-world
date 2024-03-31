package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"piaccho/cinema-api/models"
	"time"
)

func GetAllMovieReviews() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		movieId, _ := primitive.ObjectIDFromHex(c.Param("movieId"))

		// Find the movie document
		var movie models.Movie
		err := moviesCollection.FindOne(ctx, bson.M{"_id": movieId}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error finding the movie"}})
			return
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": movie.Reviews, "description": "Reviews for the movie"}},
		)
	}
}

func CreateMovieReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		movieId, _ := primitive.ObjectIDFromHex(c.Param("movieId"))

		// Find the movie document
		var movie models.Movie
		err := moviesCollection.FindOne(ctx, bson.M{"_id": movieId}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error finding the movie"}})
			return
		}

		// Bind the request body to a new Review
		var review models.Review
		if err := c.BindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error binding the request body"}})
			return
		}

		// Find the user document
		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"_id": review.UserId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error finding the user"}})
			return
		}

		// Check if the firstName matches
		if user.Firstname != review.Firstname {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"description": "The firstName does not match the user"}})
			return
		}

		// Add the new review to the movie's reviews
		review.Id = primitive.NewObjectID()
		movie.Reviews = append(movie.Reviews, review)

		// Update the movie document in the database
		_, err = moviesCollection.UpdateOne(ctx, bson.M{"_id": movieId}, bson.M{"$set": movie})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error updating the movie"}})
			return
		}

		// Return the new review
		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": review, "description": "Review successfully created"}},
		)
	}
}

func GetMovieReviewById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		movieId, _ := primitive.ObjectIDFromHex(c.Param("movieId"))
		reviewId, _ := primitive.ObjectIDFromHex(c.Param("reviewId"))

		// Find the movie document
		var movie models.Movie
		err := moviesCollection.FindOne(ctx, bson.M{"_id": movieId}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error finding the movie"}})
			return
		}

		// Find the review in the movie's reviews
		var review models.Review
		for _, r := range movie.Reviews {
			if r.Id == reviewId {
				review = r
				break
			}
		}

		// If the review was not found, return an error
		if review.Id.IsZero() {
			c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"description": "Review not found"}})
			return
		}

		// Return the review
		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": review, "description": "Review for the movie"}},
		)
	}
}

func UpdateMovieReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		movieId, _ := primitive.ObjectIDFromHex(c.Param("movieId"))
		reviewId, _ := primitive.ObjectIDFromHex(c.Param("reviewId"))

		// Find the movie document
		var movie models.Movie
		err := moviesCollection.FindOne(ctx, bson.M{"_id": movieId}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error finding the movie"}})
			return
		}

		// Find the review in the movie's reviews
		for i, r := range movie.Reviews {
			if r.Id == reviewId {
				// Bind the request body to the review
				if err := c.BindJSON(&movie.Reviews[i]); err != nil {
					c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error binding the request body"}})
					return
				}
				// Find the user document
				var user models.User
				err = userCollection.FindOne(ctx, bson.M{"_id": movie.Reviews[i].UserId}).Decode(&user)
				if err != nil {
					c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error finding the user"}})
					return
				}

				// Check if the firstName matches
				if user.Firstname != movie.Reviews[i].Firstname {
					c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"description": "The firstName does not match the user"}})
					return
				}

				break
			}
		}

		// Update the movie document in the database
		_, err = moviesCollection.UpdateOne(ctx, bson.M{"_id": movieId}, bson.M{"$set": movie})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error updating the movie"}})
			return
		}

		// Return the updated review
		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"description": "Review successfully updated"}},
		)
	}
}

func DeleteMovieReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		movieId, _ := primitive.ObjectIDFromHex(c.Param("movieId"))
		reviewId, _ := primitive.ObjectIDFromHex(c.Param("reviewId"))

		// Find the movie document
		var movie models.Movie
		err := moviesCollection.FindOne(ctx, bson.M{"_id": movieId}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error finding the movie"}})
			return
		}

		// Find the review in the movie's reviews and remove it
		for i, r := range movie.Reviews {
			if r.Id == reviewId {
				movie.Reviews = append(movie.Reviews[:i], movie.Reviews[i+1:]...)
				break
			}
		}

		// Update the movie document in the database
		_, err = moviesCollection.UpdateOne(ctx, bson.M{"_id": movieId}, bson.M{"$set": movie})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error updating the movie"}})
			return
		}

		// Return success message
		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"description": "Review successfully deleted"}},
		)
	}
}

func GetAllMovieReviewsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId, _ := primitive.ObjectIDFromHex(c.Param("userId"))

		// Find all movie documents with reviews by the user
		cursor, err := moviesCollection.Find(ctx, bson.M{"reviews.userId": userId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error finding the movies"}})
			return
		}

		// Iterate through the cursor and collect the reviews
		var reviews []models.Review
		for cursor.Next(ctx) {
			var movie models.Movie
			if err := cursor.Decode(&movie); err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error decoding the movie"}})
				return
			}
			for _, r := range movie.Reviews {
				if r.UserId == userId {
					reviews = append(reviews, r)
				}
			}
		}

		// Return the reviews
		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": reviews, "description": "Reviews by the user"}},
		)
	}
}
