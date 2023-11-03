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
	fmt.Printf("Initting")
	uri := os.Getenv("MONGODB_UFO_URI")
	if uri == "" {
		log.Fatal("Set the Mongodb uri for ufo database!")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	collation := options.Collation{
		Locale:   "en",
		Strength: 2, // 2 for case-insensitive, 1 for case-sensitive
	}
	findOptions := options.Find()
	findOptions.SetCollation(&collation)
	sightingDatabase = client.Database("i_want_to_believe")
}

func GetSightings(c *gin.Context) {
	var sightingsCollection = sightingDatabase.Collection("sightings_with_coords")
	var results []structs.Sighting

	cursor, err := sightingsCollection.Find(context.TODO(), bson.D{{"country", "finland"}}, defaultFindOptions)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Printf("No document was found with the title %s\n", "Finland")
		return
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Printf("No document was found with the country %s\n", "Finland")
		return
	}
	fmt.Printf("%s", results[0].ReportDate)
	c.IndentedJSON(http.StatusOK, results)
}

func SearchSightings(c *gin.Context) {
	var results []structs.Sighting
	var sightingsCollection = sightingDatabase.Collection("sightings_with_coords")
	searchWord := c.DefaultQuery("searchTerm", "")
	// startDate := c.DefaultQuery("startDate", "")
	// endDate := c.DefaultQuery("endDate", "")
	if len(searchWord) > 0 {
		fullTextBson := bson.D{{ "$search", searchWord}}
		searchBson := bson.D{{ "$text", fullTextBson}}
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
	}
	c.IndentedJSON(http.StatusOK, results)
}