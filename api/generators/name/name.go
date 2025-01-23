package name

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	// Register name-related questions
	generators.RegisterQuestionType("Name", GenerateNameQuestion, GenerateNameHouseQuestion)
}

// GenerateNameQuestion generates a question about a character's patronus
func GenerateNameQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with patronus
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Name")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a patronus available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate incorrect options from characters with a different name
	incorrectOptions := helpers.FilterCharactersWithDifferentAttribute(validCharacters, "Name", character.Name)

	// Ensure there are enough wrong options
	if len(incorrectOptions) < 3 {
		return nil, fmt.Errorf("not enough diverse characters with different patronus")
	}

	// Select up to 3 incorrect options
	rng.Shuffle(len(incorrectOptions), func(i, j int) { incorrectOptions[i], incorrectOptions[j] = incorrectOptions[j], incorrectOptions[i] })
	selectedIncorrect := []string{}
	for i := 0; i < len(incorrectOptions) && i < 3; i++ {
		selectedIncorrect = append(selectedIncorrect, incorrectOptions[i].Name)
	}

	// Add the correct answer to options
	options := append(selectedIncorrect, character.Name)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	// Validate options length
	if len(options) < 2 {
		return nil, fmt.Errorf("failed to generate enough options for the question")
	}

	// Generate the question
	question := fmt.Sprintf("What is the name of the character whose patronus is %s?", character.Patronus)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "hard",
		"type":       "name",
	}, nil
}

// GenerateNameHouseQuestion generates a question about which character belongs to a specific house
func GenerateNameHouseQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with a house
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "House")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a house available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate incorrect options from characters with a different name
	incorrectOptions := helpers.FilterCharactersWithDifferentAttribute(validCharacters, "Name", character.Name)

	// Ensure there are enough wrong options
	if len(incorrectOptions) < 3 {
		return nil, fmt.Errorf("not enough diverse characters with different houses")
	}

	// Generate options
	options := helpers.GenerateOptionsFromCharacters(validCharacters, character.Name, "Name", rng)

	// Generate the question
	question := fmt.Sprintf("Who belongs to the %s house?", character.House)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "easy",
		"type":       "name",
	}, nil
}
