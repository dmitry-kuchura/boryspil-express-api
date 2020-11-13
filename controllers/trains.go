package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	u "boryspil-express-api/utils"
)

type Trains struct {
	Trains []Train `json:"trains"`
}

type Train struct {
	Number              int        `json:"number"`
	DepartureTrafficHub TrafficHub `json:"departureTrafficHub"`
	ArrivalTrafficHub   TrafficHub `json:"arrivalTrafficHub"`
	Segments            []Segment  `json:"segments"`
	DepartureTime       string     `json:"departureTime"`
	ArrivalTime         string     `json:"arrivalTime"`
}

type Segment struct {
	DepartureTime string     `json:"departureTime"`
	TrafficHub    TrafficHub `json:"trafficHub"`
}

type TrafficHub struct {
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

func GetTrains(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./data.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var trains Trains

	json.Unmarshal(byteValue, &trains)

	resp := u.Message(true, "success")
	resp["data"] = trains
	u.Respond(w, resp)
}
