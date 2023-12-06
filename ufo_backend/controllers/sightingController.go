package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"ufo_backend/structs"
)

var sightingDatabase *mongo.Database
var defaultFindOptions *options.FindOptions

func init() {
	fmt.Println("Initting")
	uri := os.Getenv("MONGODB_UFO_URI")
	if uri == "" {
		log.Fatal("Set the Mongodb uri for ufo database!")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Connection failed!")
	}
	collation := options.Collation{
		Locale:   "en",
		Strength: 2,
	}
	findOptions := options.Find()
	findOptions.SetCollation(&collation)
	sightingDatabase = client.Database("i_want_to_believe")
}

type SearchTerms struct {
	SearchTerm bson.M `bson:"$search,omitempty"`
	StartDate  string `bson:"startDate,omitempty"`
	EndDate    string `bson:"endDate,omitempty"`
}

func TestEndpoint(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bson.D{{"its", "working"}})
}

func SearchSightings(c *gin.Context) {
	fmt.Println("Trying to get database connection")

	var results []structs.Sighting
	var sightingsCollection = sightingDatabase.Collection("sightings_with_coords")
	startDate := c.DefaultQuery("startDate", "")
	endDate := c.DefaultQuery("endDate", "")
	searchTerm := c.DefaultQuery("searchTerm", "")
	searchBson := bson.D{{"$text", bson.D{{"$search", searchTerm}}}}
	dateFilter := bson.M{}

	if len(startDate) > 0 {
		dateFilter["$gte"] = startDate
	}
	if len(endDate) > 0 {
		dateFilter["$lte"] = endDate
	}
	if len(startDate) > 0 || len(endDate) > 0 {
		searchBson = append(searchBson, bson.E{Key: "report_date", Value: dateFilter})
	}

	cursor, err := sightingsCollection.Find(context.TODO(), searchBson, defaultFindOptions)
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
