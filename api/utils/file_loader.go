package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadJSON(filePath string, target interface{}) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load file: %s", err)
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}
}
