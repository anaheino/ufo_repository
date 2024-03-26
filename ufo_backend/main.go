package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"ufo_backend/controllers"
)

func main() {
	router := gin.Default()
	router.Use(corsMiddleware())

	sightingController := controllers.NewSightingsController(nil)
	sightingController.Init()

	machineLearningController := controllers.NewMachineLearningController(nil)
	machineLearningController.Init()

	router.GET("/search", sightingController.SearchSightings)
	router.POST("/probability", machineLearningController.RandomForestForCoordinates)

	uri := fmt.Sprintf("%s:8080", os.Getenv("localhost"))
	fmt.Printf(uri)
	router.Run(uri)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests (OPTIONS)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
