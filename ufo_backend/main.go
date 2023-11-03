package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ufo_backend/controllers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/search", controllers.SearchSightings)
	router.Run("localhost:8080")
}