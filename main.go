package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"boryspil-express-api/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning .env file load error", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/trains", api.GetTrains).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
