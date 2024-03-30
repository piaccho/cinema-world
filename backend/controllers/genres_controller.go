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

var genreCollection *mongo.Collection = configs.GetCollection("genres")

func GetAllGenres() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Find all documents in the genreCollection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var genres []models.Genre
		defer cancel()

		results, err := genreCollection.Find(ctx, bson.M{})
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
			var singleGenre models.Genre
			if err = results.Decode(&singleGenre); err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			genres = append(genres, singleGenre)
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": genres}},
		)
	}
}

func CreateGenre() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var genre models.Genre
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&genre); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&genre); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newGenre := models.Genre{
			Id:   primitive.NewObjectID(),
			Name: genre.Name,
		}
		// Insert the new genre into the database
		result, err := genreCollection.InsertOne(ctx, newGenre)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, models.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetGenreById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		genreId := c.Param("genreId")
		var genre models.Genre
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(genreId)

		err := genreCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&genre)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No genre found with the provided ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": genre}})
	}
}

func UpdateGenre() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		genreId := c.Param("genreId")
		var genre models.Genre
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(genreId)

		// validate the request body
		if err := c.BindJSON(&genre); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&genre); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// update the genre in the database
		update := bson.M{
			"$set": bson.M{
				"name": genre.Name,
			},
		}
		result, err := genreCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// get updated user details
		var updatedGenre models.Genre
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedGenre)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedGenre}})
	}
}

func DeleteGenre() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		genreId := c.Param("genreId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(genreId)

		// delete the genre from the database
		_, err := genreCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Genre deleted successfully"}})
	}
}
