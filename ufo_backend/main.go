package main


import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

type Sighting struct {
	ID string `json:"_id"`
	Date string `json:"date"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	Shape string `json:"shape"`
	Duration string `json:"duration"`
	Description string `json:"description"`
	ReportDate string `json:"report_date"`
	HasImages bool `json:"has_images"`
	Link string `json:"link"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
var sightingsCollection *mongo.Collection

func getSightings(c *gin.Context) {
	var results []Sighting
	cursor, err := sightingsCollection.Find(context.TODO(), bson.D{{"country", "Finland"}})
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Printf("No document was found with the title %s\n", "Finland")
		return
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Printf("No document was found with the country %s\n", "Finland")
		return
	}
	c.IndentedJSON(http.StatusOK, results)
}

func main() {
	if err:= godotenv.Load();err != nil {
		log.Println("No .env file")
	}
	uri := os.Getenv("MONGODB_UFO_URI")
	if uri == "" {
		log.Fatal("Set the Mongodb uri for ufo database!")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	sightingsCollection = client.Database("i_want_to_believe").Collection("sightings_with_coords")
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/sightings", getSightings)
	router.Run("localhost:8080")
}