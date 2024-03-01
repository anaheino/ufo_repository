package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"ufo_backend/structs"
)

var randomForestURL = ""

func init() {
	pythonBackEndURL := fmt.Sprintf("%s", os.Getenv("MACHINE_LEARNING_URI"))
	if pythonBackEndURL == "" {
		pythonBackEndURL = "http://localhost:5000"
	}
	randomForestURL = pythonBackEndURL + "/coordinates/likelihood"
}

func RandomForestForCoordinates(c *gin.Context) {
	fmt.Println("Trying to get a probability")
	fmt.Println(randomForestURL)
	var coordinates structs.Coordinates
	if err := c.ShouldBindJSON(&coordinates); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	jsonData, err := json.Marshal(coordinates)
	client := &http.Client{}
	req, err := http.NewRequest("POST", randomForestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("External service returned status code %d", resp.StatusCode)})
		return
	}
	jsonData, err = io.ReadAll(resp.Body)
	var requestBody structs.Probability
	if err := json.Unmarshal(jsonData, &requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	response := structs.CoordinatesProbability{
		Coordinates: structs.Coordinates{
			Latitude:  coordinates.Latitude,
			Longitude: coordinates.Longitude,
		},
		Probability: requestBody,
	}
	c.JSON(http.StatusOK, response)
}
