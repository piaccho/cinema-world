package main

import (
	"context"
	"log"
	"piaccho/cinema-api/configs"

	"github.com/gin-gonic/gin"

	"piaccho/cinema-api/routes"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Connect to MongoDB
	err := configs.GetMongoClient()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = configs.DB.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Initialize indexes
	configs.InitializeIndexes()

	// Setup routes
	routes.DefaultRoute(router.Group("/api"))
	routes.AuthRoute(router.Group("/api/auth"))
	routes.GenreRoute(router.Group("/api/genres"))
	routes.HallRoute(router.Group("/api/halls"))
	routes.ShowingRoute(router.Group("/api/showings"))
	// TODO
	routes.MovieRoute(router.Group("/api/movies"))
	routes.UserRoute(router.Group("/api/users"))

	// Start the server
	router.Run(":" + configs.EnvPort())
}
