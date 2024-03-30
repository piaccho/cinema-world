package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"piaccho/cinema-api/configs"
	"piaccho/cinema-api/models"
	"piaccho/cinema-api/utils"
)

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginRequest
		var user models.User

		// validate the request body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// find the user by email
		userCollection := configs.GetCollection("users")
		err := userCollection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ResponseModel{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Invalid email or password"}})
			return
		}

		// compare the password
		err = utils.CompareHashedPassword(user.Password, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ResponseModel{Status: http.StatusUnauthorized, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Invalid email or password"}})
			return
		}

		// generate the token
		tokenString, err := utils.CreateToken(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error while generating the token"}})
			return
		}

		c.JSON(http.StatusOK, models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"token": tokenString}})
	}
}

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.RegisterRequest
		var user models.User

		// validate the request body
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseModel{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// check if the user already exists
		userCollection := configs.GetCollection("users")
		err := userCollection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&user)
		if err == nil {
			c.JSON(http.StatusConflict, models.ResponseModel{Status: http.StatusConflict, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "User already exists"}})
			return
		}

		// hash the password
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error while hashing the password"}})
			return
		}

		// create the new user
		newUser := models.User{
			Type:      "User",
			Email:     req.Email,
			Password:  hashedPassword,
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
		}
		_, err = userCollection.InsertOne(context.TODO(), newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ResponseModel{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error(), "description": "Error while creating the user"}})
			return
		}

		// return success message
		c.JSON(http.StatusOK, models.ResponseModel{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"description": "User successfully registered"}})
	}
}
