package ancestry

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	// Register ancestry-related questions
	generators.RegisterQuestionType("Ancestry", GenerateAncestryOfNameQuestion, GenerateWhoIsAncestryQuestion)
}

// GenerateAncestryOfNameQuestion generates a question about the ancestry of a specific character
func GenerateAncestryOfNameQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with an ancestry
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Ancestry")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with an ancestry available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate question and options
	question := fmt.Sprintf("What is the ancestry of %s?", character.Name)
	ancestries := helpers.ExtractUniqueValues(validCharacters, "Ancestry")

	// Ensure enough options
	if len(ancestries) < 4 {
		return nil, fmt.Errorf("not enough unique ancestries to generate options")
	}

	options := helpers.GenerateOptions(ancestries, character.Ancestry, rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Ancestry,
		"difficulty": "medium",
		"type":       "ancestry",
	}, nil
}

// GenerateWhoIsAncestryQuestion generates a question about which character has a specific ancestry
func GenerateWhoIsAncestryQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with an ancestry
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Ancestry")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with an ancestry available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate incorrect options from characters with different ancestry
	incorrectOptions := helpers.FilterCharactersWithDifferentAttribute(validCharacters, "Ancestry", character.Ancestry)

	// Ensure there are enough wrong options
	if len(incorrectOptions) < 3 {
		return nil, fmt.Errorf("not enough characters with different ancestry to generate options")
	}

	// Select up to 3 incorrect options
	rng.Shuffle(len(incorrectOptions), func(i, j int) { incorrectOptions[i], incorrectOptions[j] = incorrectOptions[j], incorrectOptions[i] })
	incorrectNames := []string{incorrectOptions[0].Name, incorrectOptions[1].Name, incorrectOptions[2].Name}

	// Add the correct answer
	options := append(incorrectNames, character.Name)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	// Generate the question
	question := fmt.Sprintf("Who is %s among these characters?", character.Ancestry)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "medium",
		"type":       "ancestry",
	}, nil
}
