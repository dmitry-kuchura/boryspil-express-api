package models

type Train struct {
	Number              int        `json:"number"`
	DepartureTrafficHub TrafficHub `json:"departureTrafficHub"`
	ArrivalTrafficHub   TrafficHub `json:"arrivalTrafficHub"`
	Segments            []Segment  `json:"segments"`
	DepartureTime       string     `json:"departureTime"`
	ArrivalTime         string     `json:"arrivalTime"`
}
