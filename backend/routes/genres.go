package routes

import (
	"piaccho/cinema-api/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupGenresRoutes(router *gin.RouterGroup, client *mongo.Client) {
	router.GET("/", handlers.GetGenres(client))
	// router.POST("/", handlers.CreateGenre(client))
	// router.GET("/:id", handlers.GetGenreById(client))
	// router.PUT("/:id", handlers.UpdateGenre(client))
	// router.DELETE("/:id", handlers.DeleteGenre(client))
}
