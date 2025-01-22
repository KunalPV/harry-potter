package generators

import (
	"fmt"
	"math/rand"
	"time"

	"harry-potter-api/generators/name"
	"harry-potter-api/models"
)

// GenerateQuestion serves as the entry point for generating a random trivia question
func GenerateQuestion(characters []models.Character, spells []models.Spell) (map[string]interface{}, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// List of question types
	questionTypes := []string{"Name"} // Add more types as we implement them
	selectedType := questionTypes[rng.Intn(len(questionTypes))]

	// Delegate to the appropriate generator based on question type
	switch selectedType {
	case "Name":
		if rng.Intn(2) == 0 {
			return name.GenerateNameQuestion(characters, rng)
		}
		return name.GenerateNameHouseQuestion(characters, rng)
	default:
		return nil, fmt.Errorf("unknown question type: %s", selectedType)
	}
}
