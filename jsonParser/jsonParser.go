package jsonParser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jzarzeczny/go-checker/interfaces"
)

func Parser() ([]interfaces.URLData, error) {
	file, err := os.Open("data.json")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()

	var urls []interfaces.URLData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&urls)

	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return nil, err
	}

	return urls, nil
}
