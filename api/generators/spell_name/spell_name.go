package spell_name

import (
	"fmt"
	"math/rand"

	"harry-potter-api/generators"
	"harry-potter-api/generators/helpers"
	"harry-potter-api/models"
)

func init() {
	generators.RegisterQuestionType("Spell Name", GenerateSpellNameQuestion)
}

// GenerateSpellNameQuestion generates a question about the name of a spell based on its description
func GenerateSpellNameQuestion(characters []models.Character, spells []models.Spell, rng *rand.Rand) (map[string]interface{}, error) {
	// Ensure there are valid spells
	if len(spells) == 0 {
		return nil, fmt.Errorf("no spells available to generate questions")
	}

	// Select a random spell
	spell := spells[rng.Intn(len(spells))]

	// Generate question and options
	question := fmt.Sprintf("Pick the right spell name which matches the spell description. \n Spell Description: \"%s\"", spell.Description)
	spellNames := helpers.ExtractUniqueValuesFromSpells(spells, "Name")
	options := helpers.GenerateOptions(spellNames, spell.Name, rng)

	return map[string]interface{}{
		"question":   question,
		"options":    options,
		"answer":     spell.Name,
		"difficulty": "easy",
		"type":       "spell-name",
	}, nil
}
