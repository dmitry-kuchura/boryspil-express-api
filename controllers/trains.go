package api

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	m "boryspil-express-api/models"
	u "boryspil-express-api/utils"
)

type Trains m.Trains
type Train = m.Train

func GetTrains(w http.ResponseWriter, r *http.Request) {
	byteValue, _ := u.OpenFile("./database/outbound.json")

	var trains Trains

	json.Unmarshal(byteValue, &trains)

	resp := u.Message(true, "success")
	resp["data"] = trains
	u.Respond(w, resp)
}

func GetCurrentTrains(w http.ResponseWriter, r *http.Request) {
	byteValue, _ := u.OpenFile("./database/outbound.json")

	var trains Trains
	var currentTrains []Train

	json.Unmarshal(byteValue, &trains)

	location, _ := time.LoadLocation("Europe/Kiev")
	currentTime := time.Now().In(location)

	var data = trains.Trains

	for _, train := range data {
		trainTimeDeparture := u.GetTrainTimeDeparture(train)
		trainTimeArrival := u.GetTrainTimeArrival(train)

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
