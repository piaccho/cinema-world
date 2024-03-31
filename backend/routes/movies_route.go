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
	router.GET("/ref", controllers.GetMoviesRefs())
	router.POST("/", controllers.CreateMovie())
	router.GET("/:movieId", controllers.GetMovieById())
	router.PUT("/:movieId", controllers.UpdateMovie())
	router.DELETE("/:movieId", controllers.DeleteMovie())
	// Routes for reviews
	router.POST("/:movieId/review", controllers.CreateMovieReview())
	router.GET("/:movieId/review", controllers.GetAllMovieReviews())
	router.GET("/:movieId/review/:reviewId", controllers.GetMovieReviewById())
	router.PUT("/:movieId/review/:reviewId", controllers.UpdateMovieReview())
	router.DELETE("/:movieId/review/:reviewId", controllers.DeleteMovieReview())
}
