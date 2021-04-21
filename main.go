package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	c "boryspil-express-api/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning .env file load error", err)
	}

	router := mux.NewRouter()
	subRouters := router.PathPrefix("/api").Subrouter()

	subRouters.HandleFunc("/trains", c.GetTrains).Methods("POST")
	subRouters.HandleFunc("/current", c.GetCurrentTrains).Methods("GET")
	subRouters.HandleFunc("/train/{number:[0-9]+}", c.GetTrainInfo).Methods("GET")
	subRouters.HandleFunc("/traffic-hub/search/{traffic-hub}", c.GetStationInfo).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println("Server is listening...")
	fmt.Println(port)

	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Print(err)
	}
}
