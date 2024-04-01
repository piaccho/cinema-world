package routes

import (
	"piaccho/cinema-api/controllers"

	"github.com/gin-gonic/gin"
)

func MovieRoute(router *gin.RouterGroup) {
	// CRUD operations for movies
	router.GET("/", controllers.GetMovies("all"))
	router.POST("/", controllers.CreateMovie())
	router.GET("/:movieId", controllers.GetMovieById())
	router.PUT("/:movieId", controllers.UpdateMovie())
	router.DELETE("/:movieId", controllers.DeleteMovie())
	// Routes for getting movies
	router.GET("/title/:movieTitle", controllers.GetMovieByTitle())
	router.GET("/popular/:quantity", controllers.GetMovies("popular"))
	router.GET("/upcoming/:quantity", controllers.GetMovies("upcoming"))
	router.GET("/genre/name/:genreName", controllers.GetMovies("genreName"))
	router.GET("/genre/:genreId", controllers.GetMovies("genreId"))
	router.GET("/search/:query", controllers.GetMovies("search"))
	router.GET("/ref", controllers.GetMoviesRefs())
	// Routes for CRUD reviews
	router.POST("/:movieId/reviews", controllers.CreateMovieReview())
	router.GET("/:movieId/reviews", controllers.GetAllMovieReviews())
	router.GET("/:movieId/reviews/:reviewId", controllers.GetMovieReviewById())
	router.PUT("/:movieId/reviews/:reviewId", controllers.UpdateMovieReview())
	router.DELETE("/:movieId/reviews/:reviewId", controllers.DeleteMovieReview())
}
