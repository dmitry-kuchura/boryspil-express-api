package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func OpenFile(file string) ([]byte, error) {
	jsonFile, err := os.Open("./database/outbound.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened: " + file)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	return byteValue, err
}
