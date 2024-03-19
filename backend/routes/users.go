package routes

import (
	"piaccho/cinema-api/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupUsersRoutes(router *gin.RouterGroup, client *mongo.Client) {
	router.POST("/login", handlers.LoginUser(client))
	router.POST("/register", handlers.RegisterUser(client))
}
