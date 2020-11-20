package models

type Segment struct {
	ArrivalTime   *string    `json:"arrivalTime"`
	DepartureTime *string    `json:"departureTime"`
	TrafficHub    TrafficHub `json:"trafficHub"`
}
