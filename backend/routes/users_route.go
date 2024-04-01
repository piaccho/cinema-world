package routes

import (
	"github.com/gin-gonic/gin"
	"piaccho/cinema-api/controllers"
)

func UserRoute(router *gin.RouterGroup) {
	// CRUD operations
	router.POST("/", controllers.CreateUser())
	router.GET("/:userId", controllers.GetUserById())
	router.PUT("/:userId", controllers.UpdateUser())
	router.DELETE("/:userId", controllers.DeleteUser())
	router.GET("/", controllers.GetAllUsers())
	// Routes for watchlist
	router.POST("/:userId/toWatchList/items", controllers.AddToWatchListItem())
	router.GET("/:userId/toWatchList/items", controllers.GetAllToWatchListItems())
	router.GET("/:userId/toWatchList/items/:itemId", controllers.GetToWatchListItemById())
	router.PUT("/:userId/toWatchList/items/:itemId", controllers.UpdateToWatchListItem())
	router.DELETE("/:userId/toWatchList/items/:itemId", controllers.DeleteToWatchListItem())
	// Other routes
	router.GET("/:userId/reviews/", controllers.GetAllMovieReviewsByUserId())
}
