package wand_core

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	// Register wand core-related questions
	generators.RegisterQuestionType("Wand Core", GenerateWandCoreOfNameQuestion, GenerateWhoHasWandCoreQuestion)
}

// GenerateWandCoreOfNameQuestion generates a question about the wand core of a specific character
func GenerateWandCoreOfNameQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with a wand core
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Wand.Core")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a wand core available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Extract unique wand cores
	wandCores := helpers.ExtractUniqueValues(validCharacters, "Wand.Core")

	// Ensure enough options
	if len(wandCores) < 4 {
		return nil, fmt.Errorf("not enough unique wand cores to generate options")
	}

	// Generate the question and options
	question := fmt.Sprintf("What is the core of %s's wand?", character.Name)
	options := helpers.GenerateOptions(wandCores, character.Wand.Core, rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Wand.Core,
		"difficulty": "medium",
		"type":       "wand-core",
	}, nil
}

// GenerateWhoHasWandCoreQuestion generates a question about which character has a wand with a specific core
func GenerateWhoHasWandCoreQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with a wand core
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Wand.Core")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a wand core available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate incorrect options from characters with different wand cores
	incorrectOptions := helpers.FilterCharactersWithDifferentAttribute(validCharacters, "Wand.Core", character.Wand.Core)

	// Ensure there are enough wrong options
	if len(incorrectOptions) < 3 {
		return nil, fmt.Errorf("not enough characters with different wand cores to generate options")
	}

	// Select up to 3 incorrect options
	rng.Shuffle(len(incorrectOptions), func(i, j int) { incorrectOptions[i], incorrectOptions[j] = incorrectOptions[j], incorrectOptions[i] })
	incorrectNames := []string{incorrectOptions[0].Name, incorrectOptions[1].Name, incorrectOptions[2].Name}

	// Add the correct answer to the options
	options := append(incorrectNames, character.Name)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	// Generate the question
	question := fmt.Sprintf("Whose wand core is %s?", character.Wand.Core)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "medium",
		"type":       "wand-core",
	}, nil
}
