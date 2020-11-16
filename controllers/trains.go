package api

import (
	"boryspil-express-api/models"
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
