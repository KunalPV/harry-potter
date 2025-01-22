package house

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	// Register house-related questions
	generators.RegisterQuestionType("House", GenerateHouseOfNameQuestion, GenerateMemberOfHouseQuestion)
}

// GenerateHouseOfNameQuestion generates a question about the house of a specific character
func GenerateHouseOfNameQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with a house
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "House")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a house available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate the question and options
	question := fmt.Sprintf("What is the house of %s?", character.Name)
	houses := helpers.ExtractUniqueValues(validCharacters, "House")
	options := helpers.GenerateOptions(houses, character.House, rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.House,
		"difficulty": "easy",
		"type":       "house",
	}, nil
}

// GenerateMemberOfHouseQuestion generates a question about which character is a member of a specific house
func GenerateMemberOfHouseQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with a house
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "House")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a house available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate incorrect options
	incorrectOptions := helpers.FilterCharactersWithDifferentAttribute(validCharacters, "House", character.House)
	if len(incorrectOptions) < 3 {
		return nil, fmt.Errorf("not enough characters with different houses to generate options")
	}

	// Select up to 3 incorrect options
	rng.Shuffle(len(incorrectOptions), func(i, j int) { incorrectOptions[i], incorrectOptions[j] = incorrectOptions[j], incorrectOptions[i] })
	incorrectNames := []string{incorrectOptions[0].Name, incorrectOptions[1].Name, incorrectOptions[2].Name}

	// Add the correct answer to the options
	options := append(incorrectNames, character.Name)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	// Validate options length
	if len(options) != 4 {
		return nil, fmt.Errorf("failed to generate enough unique options")
	}

	// Generate the question and options
	question := fmt.Sprintf("Which character is a member of the %s house?", character.House)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "easy",
		"type":       "house",
	}, nil
}
