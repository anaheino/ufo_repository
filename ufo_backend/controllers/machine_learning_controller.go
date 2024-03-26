package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ufo_backend/services"
	"ufo_backend/structs"
)

type MachineLearningController struct {
	machineLearningService *services.MachineLearningService
}

func NewMachineLearningController(machineLearningService *services.MachineLearningService) *MachineLearningController {
	return &MachineLearningController{machineLearningService: machineLearningService}
}
func (controller *MachineLearningController) Init() {
	controller.machineLearningService = services.NewMachineLearningService("")
	if err := controller.machineLearningService.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}

func (controller *MachineLearningController) RandomForestForCoordinates(c *gin.Context) {
	fmt.Println("Trying to get a probability")
	var coordinates structs.Coordinates
	if err := c.ShouldBindJSON(&coordinates); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	jsonData, err := json.Marshal(coordinates)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to parse JSON"})
		return
	}
	probability := controller.machineLearningService.RandomForestCoordinates(jsonData)

	response := structs.CoordinatesProbability{
		Coordinates: structs.Coordinates{
			Latitude:  coordinates.Latitude,
			Longitude: coordinates.Longitude,
		},
		Probability: probability,
	}
	c.JSON(http.StatusOK, response)
}
