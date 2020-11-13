package api

import (
	"net/http"

	u "boryspil-express-api/utils"
)

func GetTrains(w http.ResponseWriter, r *http.Request) {

	data := "{}"
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
