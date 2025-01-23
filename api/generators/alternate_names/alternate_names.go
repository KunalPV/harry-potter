package alternate_names

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	// Register alternate_names-related questions
	generators.RegisterQuestionType("Alternate Names", GenerateAlternateNameQuestion)
}

// GenerateAlternateNameQuestion generates a question about a character's alternate name
func GenerateAlternateNameQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with non-empty alternate_names
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "AlternateNames")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with alternate names available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Select a random alternate name from the character's list
	if len(character.AlternateNames) == 0 {
		return nil, fmt.Errorf("character %s has no alternate names", character.Name)
	}
	alternateName := character.AlternateNames[rng.Intn(len(character.AlternateNames))]

	// Generate incorrect options from characters with different names
	incorrectOptions := helpers.FilterCharactersWithDifferentAttribute(validCharacters, "Name", character.Name)
	if len(incorrectOptions) < 3 {
		return nil, fmt.Errorf("not enough characters with different alternate names to generate options")
	}

	// Select up to 3 incorrect options
	rng.Shuffle(len(incorrectOptions), func(i, j int) { incorrectOptions[i], incorrectOptions[j] = incorrectOptions[j], incorrectOptions[i] })
	incorrectNames := []string{incorrectOptions[0].Name, incorrectOptions[1].Name, incorrectOptions[2].Name}

	// Add the correct answer
	options := append(incorrectNames, character.Name)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	// Generate the question
	question := fmt.Sprintf("Who is also known as '%s'?", alternateName)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "medium",
		"type":       "alternate-names",
	}, nil
}
