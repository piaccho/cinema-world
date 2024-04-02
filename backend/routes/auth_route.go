package routes

import (
	"github.com/gin-gonic/gin"
	"piaccho/cinema-api/controllers"
)

func AuthRoute(router *gin.RouterGroup) {
	router.POST("/signIn", controllers.LoginUser())
	router.POST("/signUp", controllers.RegisterUser())
}
