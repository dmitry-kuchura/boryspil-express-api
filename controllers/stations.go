package api

import (
	"github.com/gorilla/mux"
	"net/http"

	u "boryspil-express-api/utils"
)

func GetStationInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	station := vars["station"]

	resp := u.Message(true, "success")
	resp["station"] = station
	u.Respond(w, resp)
}
