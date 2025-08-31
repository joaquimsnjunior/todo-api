package router

import (
	"os"

	"github.com/gin-gonic/gin"
)



func InitializeRouter() {
	// Initialize Gin router
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run("0.0.0.0:" + port)
}