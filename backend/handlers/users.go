package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginUser(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func RegisterUser(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
