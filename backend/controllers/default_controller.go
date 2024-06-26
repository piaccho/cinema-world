package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHello Hello route
func GetHello() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("name") != "" {
			name := c.Query("name")
			c.JSON(http.StatusOK, gin.H{"message": "Hello " + name + "!"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
		}
	}
}

// GetGoodbye Goodbye with given name route
func GetGoodbye() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"message": "Goodbye " + name + "!"})
	}
}
