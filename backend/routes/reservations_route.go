package routes

import (
	"github.com/gin-gonic/gin"
	"piaccho/cinema-api/controllers"
)

func ReservationRoute(router *gin.RouterGroup) {
	// CRUD operations
	router.GET("/", controllers.GetAllReservations())
	router.GET("/:reservationId", controllers.GetReservationById())
	router.POST("/", controllers.CreateReservation())
	router.PUT("/:reservationId", controllers.UpdateReservation())
	router.DELETE("/:reservationId", controllers.DeleteReservation())
	router.GET("/user/:userId", controllers.GetReservationsByUserId())
}
