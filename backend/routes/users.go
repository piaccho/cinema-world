package routes

import (
	"piaccho/cinema-api/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupUsersRoutes(router *gin.RouterGroup, client *mongo.Client) {
	router.GET("/", handlers.GetGenres(client))
}
