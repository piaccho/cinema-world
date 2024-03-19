package routes

import (
	"piaccho/cinema-api/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupMoviesRoutes(router *gin.RouterGroup, client *mongo.Client) {
	router.GET("/", handlers.GetMovies(client, "all"))
	router.GET("/popular/:quantity", handlers.GetMovies(client, "popular"))
	router.GET("/upcoming/:quantity", handlers.GetMovies(client, "upcoming"))
	router.GET("/genres/:name", handlers.GetMovies(client, "genres"))
	router.GET("/search/:query", handlers.GetMovies(client, "search"))
	// router.POST("/", handlers.CreateMovie)
	// router.GET("/:id", handlers.GetMovieById)
	// router.PUT("/:id", handlers.UpdateMovie)
	// router.DELETE("/:id", handlers.DeleteMovie)
}
