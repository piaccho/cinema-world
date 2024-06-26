package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"piaccho/cinema-api/configs"
	"piaccho/cinema-api/models"
	"piaccho/cinema-api/utils"
	"time"
)

var userCollection = configs.GetCollection("users")

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		// Check if the email already exists (Handled by MongoDB Index)

		hashedPassword, err := utils.HashPassword(user.Password)
		newUser := models.User{
			Id:        primitive.NewObjectID(),
			Type:      user.Type,
			Email:     user.Email,
			Password:  hashedPassword,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			ToWatch:   []models.ToWatchListItem{},
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, models.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userId := c.Param("userId")
		objId, _ := primitive.ObjectIDFromHex(userId)

		var user models.User

		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "No user found with the provided ID"}})
				return
			}
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get user"}})
			return
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user, "description": "User found"}})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userId := c.Param("userId")
		objId, _ := primitive.ObjectIDFromHex(userId)
		var user models.User

		// validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Validation error"}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := utils.Validator.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error(), "description": "Validation error"}})
			return
		}

		hashedPassword, err := utils.HashPassword(user.Password)
		update := bson.M{"type": user.Type, "email": user.Email, "password": hashedPassword, "firstname": user.Firstname, "lastname": user.Lastname}
		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to update user"}})
			return
		}

		// get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get updated user"}})
				return
			}
		}

		c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser, "description": "User updated successfully"}})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to delete user"}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				models.Response{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": nil, "description": "User with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": nil, "description": "User successfully deleted!"}},
		)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get users"}})
			return
		}

		// reading from the db in an optimal way
		defer func(results *mongo.Cursor, ctx context.Context) {
			err := results.Close(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get users"}})
			}
		}(results, ctx)

		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Failed to get users"}})
			}

			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			models.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users, "description": "Users found"}},
		)
	}
}
