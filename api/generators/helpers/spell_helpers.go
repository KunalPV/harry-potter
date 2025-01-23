package helpers

import (
	"harry-potter-api/models"
)

// ExtractUniqueValuesFromSpells extracts unique values from spells for a given attribute
func ExtractUniqueValuesFromSpells(spells []models.Spell, field string) []string {
	uniqueSet := make(map[string]bool)
	for _, spell := range spells {
		switch field {
		case "Name":
			uniqueSet[spell.Name] = true
		}
	}

	values := []string{}
	for value := range uniqueSet {
		values = append(values, value)
	}
	return values
}

func ExtractUniqueSpellDescriptions(spells []models.Spell) []string {
	descriptionSet := make(map[string]bool)
	for _, spell := range spells {
		descriptionSet[spell.Description] = true
	}

	descriptions := []string{}
	for description := range descriptionSet {
		descriptions = append(descriptions, description)
	}
	return descriptions
}
