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

func TestLoadOutboundJsonFile(t *testing.T) {
	_, err := os.Open("./database/outbound.json")

	if err != nil {
		t.Errorf("Didn't try and open outbound.json file")
	}
}

func TestLoadInboundJsonFile(t *testing.T) {
	_, err := os.Open("./database/inbound.json")

	if err != nil {
		t.Errorf("Didn't try and open inbound.json file")
	}
}
