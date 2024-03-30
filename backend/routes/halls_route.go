package routes

import (
	"piaccho/cinema-api/controllers"

	"github.com/gin-gonic/gin"
)

func HallRoute(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllHalls())
	router.POST("/", controllers.CreateHall())
	router.GET("/:hallId", controllers.GetHallById())
	router.PUT("/:hallId", controllers.UpdateHall())
	router.DELETE("/:hallId", controllers.DeleteHall())
}
