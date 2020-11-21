package utils

import (
	"strconv"
	"strings"
	"time"

	m "boryspil-express-api/models"
)

func GetTrainTimeDeparture(train m.Train) time.Time {
	location, _ := time.LoadLocation("Europe/Kiev")
	currentTime := time.Now().In(location)

	departureTime := strings.Split(train.DepartureTime, ":")
	departureTimeHours, _ := strconv.Atoi(departureTime[0])
	departureTimeMinutes, _ := strconv.Atoi(departureTime[1])

	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), departureTimeHours, departureTimeMinutes, currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
}

func GetTrainTimeArrival(train m.Train) time.Time {
	location, _ := time.LoadLocation("Europe/Kiev")
	currentTime := time.Now().In(location)

	arrivalTime := strings.Split(train.ArrivalTime, ":")
	arrivalTimeHours, _ := strconv.Atoi(arrivalTime[0])
	arrivalTimeMinutes, _ := strconv.Atoi(arrivalTime[1])

	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), arrivalTimeHours, arrivalTimeMinutes, currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
}
