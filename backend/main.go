package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"piaccho/cinema-api/config"
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
	// Connect to MongoDB
	client, err := config.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Create a new router
	router := gin.Default()

	// Setup CORS middleware
	router.Use(CORSMiddleware())

	// Setup routes

	defaultRoutes := router.Group("/api")
	routes.SetupDefaultRoutes(defaultRoutes)

	genresRoutes := router.Group("/api/genres")
	routes.SetupGenresRoutes(genresRoutes, client)

	moviesRoutes := router.Group("/api/movies")
	routes.SetupMoviesRoutes(moviesRoutes, client)

	showingsRoutes := router.Group("/api/showings")
	routes.SetupShowingsRoutes(showingsRoutes, client)

	// Start the server

	router.Run(":" + os.Getenv("PORT"))

}
