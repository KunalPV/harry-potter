package name

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

// GenerateNameQuestion generates a name-related trivia question
func GenerateNameQuestion(characters []models.Character, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter out characters without a patronus
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Patronus")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a patronus available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate the question and options
	question := fmt.Sprintf("What is the name of the character whose patronus is %s?", character.Patronus)
	options := helpers.GenerateOptionsFromCharacters(validCharacters, character.Name, "Name", rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "hard",
	}, nil
}

// GenerateNameHouseQuestion generates a question about which character belongs to a specific house
func GenerateNameHouseQuestion(characters []models.Character, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with a house
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "House")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a house available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate the question and options
	question := fmt.Sprintf("Who belongs to the %s house?", character.House)
	options := helpers.GenerateOptionsFromCharacters(validCharacters, character.Name, "Name", rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "easy",
	}, nil
}
