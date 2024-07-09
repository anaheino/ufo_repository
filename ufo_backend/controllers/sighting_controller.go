package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ufo_backend/services"
	"ufo_backend/structs"
)

type SightingController struct {
	sightingService *services.SightingService
}

func NewSightingsController(sightingService *services.SightingService) *SightingController {
	return &SightingController{sightingService}
}
func (controller *SightingController) Init() {
	controller.sightingService = services.NewSightingService(nil, nil)
	if err := controller.sightingService.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}

func (controller *SightingController) SearchSightings(c *gin.Context) {
	startDate := c.DefaultQuery("startDate", "")
	endDate := c.DefaultQuery("endDate", "")
	searchTerm := c.DefaultQuery("searchTerm", "")
	results := controller.sightingService.QuerySightings(searchTerm, startDate, endDate)
	c.IndentedJSON(http.StatusOK, results)
}

func (controller *SightingController) CreateSighting(c *gin.Context) {
	var sighting structs.Sighting
	if err := c.ShouldBindJSON(&sighting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	results := controller.sightingService.CreateSighting(sighting)
	c.IndentedJSON(http.StatusOK, results)
}
