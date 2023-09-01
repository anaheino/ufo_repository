package main


import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sighting struct {
	ID     string  `json:"_id"`
	Date  string  `json:"date"`
	City string  `json:"city"`
	State  string `json:"state"`
	Country  string  `json:"country"`
	Shape string  `json:"shape"`
	Duration  string `json:"duration"`
	Description  string  `json:"description"`
	ReportDate string  `json:"report_date"`
	HasImages  bool `json:"has_images"`
	Link  string `json:"link"`
}

var sightings = []sighting{
	{ID: "1", Date:"1/1/1970T00:00:00", City: "Helsinki", State: "", Country: "Finland", Shape:"Orb", Duration: "5 minutes", Description: "Descr", ReportDate: "1/1/1970T00:00:00", HasImages: false, Link:"somelink" },
}

func getSightings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sightings)
}

func main() {
	router := gin.Default()
	router.GET("/sightings", getSightings)

	router.Run("localhost:8080")
}