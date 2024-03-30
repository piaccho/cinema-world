package routes

import (
	"github.com/gin-gonic/gin"
	"piaccho/cinema-api/controllers"
)

func AuthRoute(router *gin.RouterGroup) {
	router.POST("/login", controllers.LoginUser())
	router.POST("/register", controllers.RegisterUser())
}
