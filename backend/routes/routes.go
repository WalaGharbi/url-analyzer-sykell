package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all routes for the API
func SetupRoutes(router *gin.Engine) {
	// test check
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Test check"})
	})


	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Sykell URL Analyzer test task for Sykell"})
	})
}
