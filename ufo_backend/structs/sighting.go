package structs

type Sighting struct {
	Id string `json:"_id"`
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