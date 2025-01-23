package facts

import (
	"encoding/json"
	"log"
	"os"
)

var facts []Fact

// LoadFacts initializes facts from the facts.json file
func LoadFacts(factsFile string) {
	data, err := os.ReadFile(factsFile)
	if err != nil {
		log.Fatalf("Error loading facts file: %v", err)
	}
	if err := json.Unmarshal(data, &facts); err != nil {
		log.Fatalf("Error parsing facts file: %v", err)
	}
}
