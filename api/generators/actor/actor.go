package actor

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	// Register actor-related questions
	generators.RegisterQuestionType("Actor", GenerateActorOfCharacterQuestion, GenerateCharacterOfActorQuestion)
}

// GenerateActorOfCharacterQuestion generates a question about which actor portrayed a specific character
func GenerateActorOfCharacterQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with an actor
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Actor")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with an actor available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate question and options
	question := fmt.Sprintf("Which actor portrayed %s in the movies?", character.Name)
	actors := helpers.ExtractUniqueValues(validCharacters, "Actor")
	options := helpers.GenerateOptions(actors, character.Actor, rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Actor,
		"difficulty": "medium",
		"type":       "actor-of-character",
	}, nil
}

// GenerateCharacterOfActorQuestion generates a question about which character was played by a specific actor
func GenerateCharacterOfActorQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Filter characters with an actor
	validCharacters := helpers.FilterCharactersWithAttribute(characters, "Actor")
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with an actor available")
	}

	// Select a random character
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate incorrect options
	incorrectOptions := helpers.FilterCharactersWithDifferentAttribute(validCharacters, "Actor", character.Actor)

	// Ensure there are enough wrong options
	if len(incorrectOptions) < 3 {
		return nil, fmt.Errorf("not enough characters with different actors to generate options")
	}

	// Select up to 3 incorrect options
	rng.Shuffle(len(incorrectOptions), func(i, j int) { incorrectOptions[i], incorrectOptions[j] = incorrectOptions[j], incorrectOptions[i] })
	incorrectNames := []string{incorrectOptions[0].Name, incorrectOptions[1].Name, incorrectOptions[2].Name}

	// Add the correct answer to options
	options := append(incorrectNames, character.Name)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	// Generate the question
	question := fmt.Sprintf("Which character was played by %s?", character.Actor)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     character.Name,
		"difficulty": "medium",
		"type":       "character-of-actor",
	}, nil
}
