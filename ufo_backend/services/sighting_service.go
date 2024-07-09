package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"ufo_backend/structs"
)

type SightingService struct {
	db                 *mongo.Database
	defaultFindOptions *options.FindOptions
}

func NewSightingService(db *mongo.Database, defaultFindOptions *options.FindOptions) *SightingService {
	return &SightingService{db: db, defaultFindOptions: defaultFindOptions}
}

func (s *SightingService) Initialize() error {
	uri := os.Getenv("MONGODB_UFO_URI")
	if uri == "" {
		log.Fatal("Set the Mongodb uri for ufo database!")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf(uri)
		log.Fatal("Connection failed!")
	}
	collation := options.Collation{
		Locale:   "en",
		Strength: 2,
	}
	findOptions := options.Find()
	findOptions.SetCollation(&collation)
	s.db = client.Database("i_want_to_believe")
	s.defaultFindOptions = findOptions
	log.Printf("Mongo connection should be fine!")
	return nil
}

func (s *SightingService) QuerySightings(searchTerm string, startDate string, endDate string) []structs.Sighting {
	var sightingsCollection = s.db.Collection("sightings_with_coords")
	var results []structs.Sighting
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
	cursor, err := sightingsCollection.Find(context.TODO(), searchBson, s.defaultFindOptions)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Printf("No document was found with the word %s\n", searchTerm)
		return nil
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Printf("Found no results %s\n", searchTerm)
		return nil
	}
	return results
}

func (s *SightingService) CreateSighting(sighting structs.Sighting) string {
	var sightingsCollection = s.db.Collection("sightings_with_coords")
	insertResult, err := sightingsCollection.InsertOne(context.TODO(), sighting)
	if err != nil {
		fmt.Printf(err.Error())
		fmt.Printf("Failed to insert document: %s\n", sighting)
		return ""
	}
	return insertIDToString(insertResult.InsertedID)
}

func insertIDToString(id interface{}) string {
	switch v := id.(type) {
	case primitive.ObjectID:
		return v.Hex()
	default:
		return fmt.Sprintf("%v", v)
	}
}
