package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ufo_backend/controllers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/sightings", controllers.GetSightings)
	router.GET("/by-country", controllers.GetByCountry)
	router.Run("localhost:8080")
}