package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ufo_backend/structs"
)

var pythonBackEndURL = "http://localhost:5000"
var randomForestURL = pythonBackEndURL + "/coordinates/likelihood"

func RandomForestForCoordinates(c *gin.Context) {
	fmt.Println("Trying to get a probability")

	var coordinates structs.Coordinates
	if err := c.ShouldBindJSON(&coordinates); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	jsonData, err := json.Marshal(coordinates)
	client := &http.Client{}
	fmt.Println(jsonData)
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
	fmt.Println("trying")
	var requestBody structs.Probability
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	fmt.Println("proba")
	fmt.Println(requestBody.Probability)
	response := structs.CoordinatesProbability{
		Coordinates: structs.Coordinates{
			Latitude:  coordinates.Latitude,
			Longitude: coordinates.Longitude,
		},
		Probability: requestBody,
	}
	fmt.Println(response)
	c.JSON(http.StatusOK, response)
}
