package helpers

import (
	"math/rand"

	"harry-potter-api/models"
)

// FilterCharactersWithAttribute filters characters based on a specific attribute
func FilterCharactersWithAttribute(characters []models.Character, attribute string) []models.Character {
	filtered := []models.Character{}
	for _, char := range characters {
		switch attribute {
		case "Patronus":
			if char.Patronus != "" {
				filtered = append(filtered, char)
			}
		case "House":
			if char.House != "" {
				filtered = append(filtered, char)
			}
		}
	}
	return filtered
}

// GenerateOptionsFromCharacters generates multiple-choice options based on a field
func GenerateOptionsFromCharacters(characters []models.Character, correctAnswer string, field string, rng *rand.Rand) []string {
	options := []string{}
	for _, char := range characters {
		switch field {
		case "Name":
			if char.Name != correctAnswer {
				options = append(options, char.Name)
			}
			// Add cases for other fields as needed
		}
	}

	// Shuffle and limit options
	if len(options) > 3 {
		options = options[:3]
	}
	options = append(options, correctAnswer)
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
	return options
}
