package routes

import (
	"piaccho/cinema-api/controllers"

	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.RouterGroup) {
	router.GET("/", controllers.GetMovies("all"))
	router.GET("/popular/:quantity", controllers.GetMovies("popular"))
	router.GET("/upcoming/:quantity", controllers.GetMovies("upcoming"))
	router.GET("/genres/:name", controllers.GetMovies("genres"))
	router.GET("/search/:query", controllers.GetMovies("search"))
	router.POST("/", controllers.CreateMovie())
	router.GET("/:id", controllers.GetMovieById())
	router.PUT("/:id", controllers.UpdateMovie())
	router.DELETE("/:id", controllers.DeleteMovie())
	// Routes for reviews
	router.POST("/:id/review", controllers.CreateReview())
	router.GET("/:id/review", controllers.GetReviews())
	router.DELETE("/:id/review/:reviewId", controllers.DeleteReview())
}
