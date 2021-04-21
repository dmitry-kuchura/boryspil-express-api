package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	m "boryspil-express-api/models"
	u "boryspil-express-api/utils"
)

type TrafficHubs m.TrafficHubs
type TrafficHub m.TrafficHub

func GetStationInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	station := vars["traffic-hub"]

	byteValue, _ := u.OpenFile("./database/traffic-hubs.json")

	var stations TrafficHubs

	_ = json.Unmarshal(byteValue, &stations)

	resp := u.Message(true, "success")
	resp["data"] = stations.Trains
	resp["station"] = station
	u.Respond(w, resp)
}
