package generators

import (
	"fmt"
	"math/rand"
	"time"

	"harry-potter-api/models"
)

type QuestionGenerator func([]models.Character, []models.Spell, *rand.Rand) (map[string]interface{}, error)

var questionRegistry = map[string][]QuestionGenerator{}

// RegisterQuestionType registers a question type and its associated generators
func RegisterQuestionType(questionType string, generators ...QuestionGenerator) {
	questionRegistry[questionType] = append(questionRegistry[questionType], generators...)
}

// GenerateQuestion dynamically selects and generates a question
func GenerateQuestion(characters []models.Character, spells []models.Spell) (map[string]interface{}, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Get all question types
	types := make([]string, 0, len(questionRegistry))
	for questionType := range questionRegistry {
		types = append(types, questionType)
	}

	if len(types) == 0 {
		return nil, fmt.Errorf("no question types registered")
	}

	// Shuffle question types to randomize order
	rng.Shuffle(len(types), func(i, j int) { types[i], types[j] = types[j], types[i] })

	// Attempt each question type once
	for _, questionType := range types {
		generators := questionRegistry[questionType]

		// Shuffle generators for this type
		rng.Shuffle(len(generators), func(i, j int) { generators[i], generators[j] = generators[j], generators[i] })

		// Attempt each generator
		for _, generator := range generators {
			question, err := generator(characters, spells, rng)
			if err == nil {
				return question, nil
			}
			fmt.Printf("Failed to generate question of type %s: %s\n", questionType, err)
		}
	}

	// If no valid question could be generated
	return nil, fmt.Errorf("unable to generate a valid question after trying all types")
}
