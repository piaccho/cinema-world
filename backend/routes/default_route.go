package routes

import (
	"piaccho/cinema-api/controllers"

	"github.com/gin-gonic/gin"
)

func DefaultRoute(router *gin.RouterGroup) {
	router.GET("/", controllers.GetHello())
	router.GET("/:name", controllers.GetGoodbye())
}
