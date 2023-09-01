package main


import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

type sighting struct {
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
}
var sightings_collection = sighting{}

func getSightings(c *gin.Context) {
	var result bson.M
	err := sightings_collection.FindOne(context.TODO(), bson.D{{"city", "Turku"}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", "Turku")
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, jsonData)
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
	sightings_collection = client.Database("i_want_to_believe").Collection("sightings")


	router := gin.Default()
	router.GET("/sightings", getSightings)
	router.Run("localhost:8080")
}