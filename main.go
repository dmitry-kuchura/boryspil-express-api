package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	c "boryspil-express-api/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning .env file load error", err)
	}

	router := mux.NewRouter()
	subRouters := router.PathPrefix("/api").Subrouter()

	subRouters.HandleFunc("/trains", c.GetTrains).Methods("GET")         // DONE
	subRouters.HandleFunc("/current", c.GetCurrentTrains).Methods("GET") // DONE
	subRouters.HandleFunc("/upcoming/{station}", c.GetUpcomingTrains).Methods("GET")
	subRouters.HandleFunc("/train/{number:[0-9]+}", c.GetTrainInfo).Methods("GET") // DONE
	subRouters.HandleFunc("/station/{station}", c.GetStationInfo).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println("Server is listening...")
	fmt.Println(port)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
