package routes

import (
	"piaccho/cinema-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupDefaultRoutes(router *gin.RouterGroup) {
	router.GET("/", handlers.GetHello())
	router.GET("/:name", handlers.GetGoodbye())
}
