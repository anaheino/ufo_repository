package structs

import "go.mongodb.org/mongo-driver/bson"

type Coordinates struct {
	Latitude  interface{} `json:"latitude"`
	Longitude interface{} `json:"longitude"`
}
type Sighting struct {
	Id          string      `json:"_id"`
	Date        string      `json:"date"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Shape       string      `json:"shape"`
	Duration    string      `json:"duration"`
	Description string      `json:"description"`
	ReportDate  string      `json:"report_date"`
	HasImages   bool        `json:"has_images"`
	Link        string      `json:"link"`
	Latitude    interface{} `json:"latitude"`
	Longitude   interface{} `json:"longitude"`
}

type Probability struct {
	Probability interface{} `json:"probability"`
}

type CoordinatesProbability struct {
	Coordinates
	Probability
}
type SearchTerms struct {
	SearchTerm bson.M `bson:"$search,omitempty"`
	StartDate  string `bson:"startDate,omitempty"`
	EndDate    string `bson:"endDate,omitempty"`
}
