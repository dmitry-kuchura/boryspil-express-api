package api

import (
	"boryspil-express-api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	u "boryspil-express-api/utils"
)

type Trains struct {
	Trains []Train `json:"trains"`
}

type Train = models.Train

func GetTrains(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./database/data.json")
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

func GetCurrentTrains(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./database/data.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var trains Trains
	var train Train

	json.Unmarshal(byteValue, &trains)

	currentTime := time.Now()

	var data = trains.Trains
	var current = train

	for _, v := range data {
		trainTimeDeparture := GetTrainTimeDeparture(v)
		trainTimeArrival := GetTrainTimeArrival(v)

		diffDeparture := currentTime.Sub(trainTimeDeparture)
		diffArrival := currentTime.Sub(trainTimeArrival)

		if !math.Signbit(diffDeparture.Minutes()) && math.Signbit(diffArrival.Minutes()) {
			current = v
		}
	}

	resp := u.Message(true, "success")

	resp["data"] = current
	u.Respond(w, resp)
}

func GetTrainTimeDeparture(train Train) time.Time {
	currentTime := time.Now()

	departureTime := strings.Split(train.DepartureTime, ":")
	departureTimeHours, _ := strconv.Atoi(departureTime[0])
	departureTimeMinutes, _ := strconv.Atoi(departureTime[1])

	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), departureTimeHours, departureTimeMinutes, currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
}

func GetTrainTimeArrival(train Train) time.Time {
	currentTime := time.Now()

	arrivalTime := strings.Split(train.ArrivalTime, ":")
	arrivalTimeHours, _ := strconv.Atoi(arrivalTime[0])
	arrivalTimeMinutes, _ := strconv.Atoi(arrivalTime[1])

	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), arrivalTimeHours, arrivalTimeMinutes, currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
}
