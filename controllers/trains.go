package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
	"time"

	m "boryspil-express-api/models"
	u "boryspil-express-api/utils"
)

type Trains m.Trains
type Train = m.Train

func GetTrains(w http.ResponseWriter, _ *http.Request) {
	byteValue, _ := u.OpenFile("./database/trains.json")

	var trains Trains

	_ = json.Unmarshal(byteValue, &trains)

	resp := u.Message(true, "success")
	resp["data"] = trains.Trains
	u.Respond(w, resp)
}

func GetCurrentTrains(w http.ResponseWriter, _ *http.Request) {
	trainsData, _ := u.OpenFile("./database/trains.json")

	var trains Trains
	var currentTrains []Train

	_ = json.Unmarshal(trainsData, &trains)

	location, _ := time.LoadLocation("Europe/Kiev")
	currentTime := time.Now().In(location)

	var data = trains.Trains

	for _, train := range data {
		trainTimeDeparture := u.GetTrainTimeDeparture(train, currentTime)
		trainTimeArrival := u.GetTrainTimeArrival(train, currentTime)

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

func GetUpcomingTrains(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	station := vars["station"]

	resp := u.Message(true, "success")
	resp["station"] = station
	u.Respond(w, resp)
}

func GetTrainInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trainNumber, _ := strconv.Atoi(vars["number"])

	var trains Trains
	var currentTrain Train

	trainsData, _ := u.OpenFile("./database/trains.json")

	_ = json.Unmarshal(trainsData, &trains)

	var data = trains.Trains

	for _, train := range data {
		fmt.Println(trainNumber)
		fmt.Println(train.Number)

		if trainNumber == train.Number {
			currentTrain = train
		}
	}

	resp := u.Message(true, "success")
	resp["train"] = currentTrain
	u.Respond(w, resp)
}
