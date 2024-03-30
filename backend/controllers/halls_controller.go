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

var hallsCollection = configs.GetCollection("halls")

func GetAllHalls() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Find all documents in the hallsCollection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var halls []models.Hall
		defer cancel()

		results, err := hallsCollection.Find(ctx, bson.M{})
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
			var singleHall models.Hall
			if err = results.Decode(&singleHall); err != nil {
				c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			halls = append(halls, singleHall)
		}

		c.JSON(http.StatusOK,
			models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": halls}},
		)
	}
}

func CreateHall() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var hall models.Hall
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&hall); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&hall); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newHall := models.Hall{
			Id:          primitive.NewObjectID(),
			Name:        hall.Name,
			Rows:        hall.Rows,
			SeatsPerRow: hall.SeatsPerRow,
		}
		// Insert the new hall into the database
		result, err := hallsCollection.InsertOne(ctx, newHall)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, models.ResponseModel{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetHallById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hallId := c.Param("hallId")
		var hall models.Hall
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(hallId)

		err := hallsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&hall)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.ResponseModel{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No hall found with the provided ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": hall}})
	}
}

func UpdateHall() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hallId := c.Param("hallId")
		var hall models.Hall
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(hallId)

		// validate the request body
		if err := c.BindJSON(&hall); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&hall); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// update the hall in the database
		update := bson.M{"name": hall.Name, "rows": hall.Rows, "seats_per_row": hall.SeatsPerRow}
		result, err := hallsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// get updated user details
		var updatedHall models.Hall
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedHall)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedHall}})
	}
}

func DeleteHall() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		hallId := c.Param("hallId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(hallId)

		// delete the hall from the database
		_, err := hallsCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Hall deleted successfully"}})
	}
}
