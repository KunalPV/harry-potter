package utils

import (
	"encoding/json"
	"log"
	"os"
)

func LoadJSON(filePath string, target interface{}) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load file: %s", err)
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}
}
