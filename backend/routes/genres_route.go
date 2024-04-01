package routes

import (
	"piaccho/cinema-api/controllers"

	"github.com/gin-gonic/gin"
)

func GenreRoute(router *gin.RouterGroup) {
	router.GET("/", controllers.GetAllGenres())
	router.POST("/", controllers.CreateGenre())
	router.GET("/:genreId", controllers.GetGenreById())
	router.PUT("/:genreId", controllers.UpdateGenre())
	router.DELETE("/:genreId", controllers.DeleteGenre())

	router.GET("/name/:genreName", controllers.GetGenreByName())
}
