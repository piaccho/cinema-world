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
	//Routes for reservations
	router.POST("/:userId/reservations", controllers.CreateReservation())
	router.GET("/:userId/reservations", controllers.GetAllReservations())
	router.GET("/:userId/reservations/:reservationId", controllers.GetReservationById())
	router.DELETE("/:userId/reservations/:reservationId", controllers.DeleteReservation())
	// Routes for watchlist
	router.POST("/:userId/watchlist", controllers.AddWatchlistEntry())
	router.GET("/:userId/watchlist", controllers.GetWatchlistEntries())
	router.GET("/:userId/watchlist/:watchlistEntryId", controllers.GetWatchlistEntryById())
	router.DELETE("/:userId/watchlist/:watchlistEntryId", controllers.DeleteWatchlistEntry())
}
