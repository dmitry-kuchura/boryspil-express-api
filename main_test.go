package main

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	err := godotenv.Load()

	if err != nil {
		t.Errorf("Didn't try and open .env by default")
	}
}

func TestLoadTrainsJsonFile(t *testing.T) {
	_, err := os.Open("./database/trains.json")

	if err != nil {
		t.Errorf("Didn't try and open trains.json file")
	}
}
