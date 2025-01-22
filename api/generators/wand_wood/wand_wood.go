package wand_wood

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	// Register wand wood-related questions
	generators.RegisterQuestionType("Wand Wood", GenerateWandWoodOfNameQuestion, GenerateWhoHasWandWoodQuestion)
}

// GenerateWandWoodOfNameQuestion generates a question about the wand wood of a specific character
func GenerateWandWoodOfNameQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with a wand wood
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Wand.Wood")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a wand wood available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Extract unique wand wood types
	wandWoods := helpers.ExtractUniqueValues(validCharacters, "Wand.Wood")

	// Ensure there are enough options
	if len(wandWoods) < 4 {
		return nil, fmt.Errorf("not enough unique wand woods to generate options")
	}

	// Generate the question and options
	question := fmt.Sprintf("What is the wood type of %s's wand?", character.Name)
	options := helpers.GenerateOptions(wandWoods, character.Wand.Wood, rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Wand.Wood,
		"difficulty": "hard",
		"type":       "wand-wood",
	}, nil
}

// GenerateWhoHasWandWoodQuestion generates a question about which character has a wand made of a specific wood type
func GenerateWhoHasWandWoodQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with a wand wood
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Wand.Wood")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a wand wood available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate incorrect options from characters with different wand woods
	incorrectOptions := helpers.FilterCharactersWithDifferentAttribute(validCharacters, "Wand.Wood", character.Wand.Wood)

	// Ensure there are enough wrong options
	if len(incorrectOptions) < 3 {
		return nil, fmt.Errorf("not enough characters with different wand woods to generate options")
	}

	// Select up to 3 incorrect options
	rng.Shuffle(len(incorrectOptions), func(i, j int) { incorrectOptions[i], incorrectOptions[j] = incorrectOptions[j], incorrectOptions[i] })
	incorrectNames := []string{incorrectOptions[0].Name, incorrectOptions[1].Name, incorrectOptions[2].Name}

	// Add the correct answer to options
	options := append(incorrectNames, character.Name)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	// Generate the question
	question := fmt.Sprintf("Whose wand is made of %s?", character.Wand.Wood)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "hard",
		"type":       "wand-wood",
	}, nil
}
