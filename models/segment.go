package models

type Segment struct {
	DepartureTime string     `json:"departureTime"`
	TrafficHub    TrafficHub `json:"trafficHub"`
}
