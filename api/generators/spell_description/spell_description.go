package spell_description

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	// Register spell description-related question
	generators.RegisterQuestionType("Spell Description", GenerateSpellEffectQuestion)
}

// GenerateSpellEffectQuestion generates a question about the effect of a specific spell
func GenerateSpellEffectQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Ensure there are valid spells
	if len(spells) == 0 {
		return nil, fmt.Errorf("no spells available to generate questions")
	}

	// Select a random spell
	spell := spells[rng.Intn(len(spells))]

	// Extract unique spell descriptions
	descriptions := helpers.ExtractUniqueSpellDescriptions(spells)

	// Ensure enough options
	if len(descriptions) < 4 {
		return nil, fmt.Errorf("not enough unique spell descriptions to generate options")
	}

	// Generate question and options
	question := fmt.Sprintf(`What is the effect of the spell "%s"?`, spell.Name)
	options := helpers.GenerateOptions(descriptions, spell.Description, rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     spell.Description,
		"difficulty": "medium",
		"type":       "spell-description",
	}, nil
}
