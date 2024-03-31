package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"piaccho/cinema-api/models"
	"time"
)

func AddToWatchListItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userId, _ := primitive.ObjectIDFromHex(c.Param("userId"))

		var item models.ToWatchListItem
		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Invalid request body"}})
			return
		}

		// Check if the movie exists
		var movie models.Movie
		err := moviesCollection.FindOne(ctx, bson.M{"_id": item.Movie.Id}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Movie does not exist"}})
			return
		}

		item.Id = primitive.NewObjectID()

		// Get the user
		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Failed to get user"}})
			return
		}

		// Initialize the ToWatch list if it's null
		if user.ToWatch == nil {
			user.ToWatch = []models.ToWatchListItem{}
		}

		// Check if the movie is already in the watchlist
		for _, v := range user.ToWatch {
			if v.Movie.Id == item.Movie.Id {
				c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"message": "Movie already in watchlist"}})
				return
			}
		}

		update := bson.M{"$push": bson.M{"to_watch": item}}
		_, err = userCollection.UpdateOne(ctx, bson.M{"_id": userId}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Failed to add item to watchlist"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": item, "message": "Item added to watchlist"}})
	}
}

func GetAllToWatchListItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId, _ := primitive.ObjectIDFromHex(c.Param("userId"))
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Failed to get watchlist"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user.ToWatch, "message": "Watchlist items retrieved"}})
	}
}

func GetToWatchListItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId, _ := primitive.ObjectIDFromHex(c.Param("userId"))
		itemId, _ := primitive.ObjectIDFromHex(c.Param("itemId"))

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"_id": userId, "to_watch._id": itemId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Failed to get watchlist item"}})
			return
		}

		var item models.ToWatchListItem
		for _, v := range user.ToWatch {
			if v.Id == itemId {
				item = v
				break
			}
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": item, "message": "Watchlist item retrieved"}})
	}
}

func UpdateToWatchListItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId, _ := primitive.ObjectIDFromHex(c.Param("userId"))
		itemId, _ := primitive.ObjectIDFromHex(c.Param("itemId"))

		var item models.ToWatchListItem
		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Invalid request body"}})
			return
		}

		// Check if the movie exists
		var movie models.Movie
		err := moviesCollection.FindOne(ctx, bson.M{"_id": item.Movie.Id}).Decode(&movie)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Movie does not exist"}})
			return
		}

		// Get the user
		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Failed to get user"}})
			return
		}

		// Check if the movie is already in the watchlist
		for _, v := range user.ToWatch {
			if v.Movie.Id == item.Movie.Id {
				c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"message": "Movie already in watchlist"}})
				return
			}
		}

		update := bson.M{"$set": bson.M{"to_watch.$[elem].movie": item.Movie}}
		arrayFilter := bson.A{bson.M{"elem._id": itemId}}
		opts := options.Update().SetArrayFilters(options.ArrayFilters{Filters: arrayFilter})
		_, err = userCollection.UpdateOne(ctx, bson.M{"_id": userId}, update, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Failed to update watchlist item"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": item, "message": "Watchlist item updated"}})
	}
}

func DeleteToWatchListItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId, _ := primitive.ObjectIDFromHex(c.Param("userId"))
		itemId, _ := primitive.ObjectIDFromHex(c.Param("itemId"))

		update := bson.M{"$pull": bson.M{"to_watch": bson.M{"_id": itemId}}}
		_, err := userCollection.UpdateOne(ctx, bson.M{"_id": userId}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "message": "Failed to remove item from watchlist"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": "Item removed from watchlist"}})
	}
}
