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
	// Routes for reservations
	router.GET("/:userId/reservations", controllers.GetAllReservations())
	router.GET("/:userId/reservations/:reservationId", controllers.GetReservationById())
	router.POST("/:userId/reservations", controllers.CreateReservation())
	router.PUT("/:userId/reservations/:reservationId", controllers.UpdateReservation())
	router.DELETE("/:userId/reservations/:reservationId", controllers.DeleteReservation())
	// Routes for watchlist
	router.POST("/:userId/toWatchList", controllers.AddToWatchListItem())
	router.GET("/:userId/toWatchList", controllers.GetAllToWatchListItems())
	router.GET("/:userId/toWatchList/items/:itemId", controllers.GetToWatchListItemById())
	router.PUT("/:userId/toWatchList/items/:itemId", controllers.UpdateToWatchListItem())
	router.DELETE("/:userId/toWatchList/items/:itemId", controllers.DeleteToWatchListItem())
	// Other routes
	router.GET("/:userId/reviews/", controllers.GetAllMovieReviewsByUserId())
}
