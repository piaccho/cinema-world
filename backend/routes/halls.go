package routes

import (
	"piaccho/cinema-api/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupHallsRoutes(router *gin.RouterGroup, client *mongo.Client) {
	router.GET("/", handlers.GetHalls(client))
	// router.POST("/", handlers.CreateHall(client))
	// router.GET("/:id", handlers.GetHallById(client))
	// router.PUT("/:id", handlers.UpdateHall(client))
	// router.DELETE("/:id", handlers.DeleteHall(client))
}
