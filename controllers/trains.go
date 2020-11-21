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
	jsonFile, err := os.Open("./database/outbound.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened outbound.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var trains Trains

	json.Unmarshal(byteValue, &trains)

	resp := u.Message(true, "success")
	resp["data"] = trains
	u.Respond(w, resp)
}

func GetCurrentTrains(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./database/outbound.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened outbound.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var trains Trains
	var currentTrains []Train

	json.Unmarshal(byteValue, &trains)

	location, _ := time.LoadLocation("Europe/Kiev")
	currentTime := time.Now().In(location)

	var data = trains.Trains

	for _, train := range data {
		trainTimeDeparture := GetTrainTimeDeparture(train)
		trainTimeArrival := GetTrainTimeArrival(train)

		diffDeparture := currentTime.Sub(trainTimeDeparture)
		diffArrival := currentTime.Sub(trainTimeArrival)

		if !math.Signbit(diffDeparture.Minutes()) && math.Signbit(diffArrival.Minutes()) {
			currentTrains = append(currentTrains, train)
		}
	}

	resp := u.Message(true, "success")

	resp["data"] = currentTrains
	u.Respond(w, resp)
}

func GetTrainTimeDeparture(train Train) time.Time {
	location, _ := time.LoadLocation("Europe/Kiev")
	currentTime := time.Now().In(location)

	departureTime := strings.Split(train.DepartureTime, ":")
	departureTimeHours, _ := strconv.Atoi(departureTime[0])
	departureTimeMinutes, _ := strconv.Atoi(departureTime[1])

	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), departureTimeHours, departureTimeMinutes, currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
}

func GetTrainTimeArrival(train Train) time.Time {
	location, _ := time.LoadLocation("Europe/Kiev")
	currentTime := time.Now().In(location)

	arrivalTime := strings.Split(train.ArrivalTime, ":")
	arrivalTimeHours, _ := strconv.Atoi(arrivalTime[0])
	arrivalTimeMinutes, _ := strconv.Atoi(arrivalTime[1])

	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), arrivalTimeHours, arrivalTimeMinutes, currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
}
