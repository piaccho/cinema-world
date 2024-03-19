package routes

import (
	"piaccho/cinema-api/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupShowingsRoutes(router *gin.RouterGroup, client *mongo.Client) {
	router.GET("/date/:date", handlers.GetShowingsByDate(client))
	router.GET("/movie/:id", handlers.GetShowingsByMovieId(client))
}
