package routes

import (
	"piaccho/cinema-api/controllers"

	"github.com/gin-gonic/gin"
)

func ShowingRoute(router *gin.RouterGroup) {
	// CRUD operations
	router.GET("/", controllers.GetAllShowings())
	router.GET("/:showingId", controllers.GetShowingById())
	router.POST("/", controllers.CreateShowing())
	router.PUT("/:showingId", controllers.UpdateShowing())
	router.DELETE("/:showingId", controllers.DeleteShowing())
	// Custom queries
	router.GET("/date/:datetime", controllers.GetShowingsByDate())
	router.GET("/movie/:movieId", controllers.GetShowingsByMovieId())
	router.GET("/hall/:hallId", controllers.GetShowingsByHallId())
	// Custom operations
	router.GET("/generate/forDays/:daysNumber", controllers.GenerateShowingsForDays())
}
