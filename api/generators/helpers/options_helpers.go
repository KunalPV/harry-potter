package helpers

import (
	"harry-potter-api/models"
	"math/rand"
)

// ExtractUniqueValues extracts unique values for a specific field from the characters
func ExtractUniqueValues(characters []models.Character, field string) []string {
	uniqueSet := make(map[string]bool)
	for _, char := range characters {
		switch field {
		case "House":
			if char.House != "" {
				uniqueSet[char.House] = true
			}
		case "Ancestry":
			if char.Ancestry != "" {
				uniqueSet[char.Ancestry] = true
			}
		case "Wand.Wood":
			if char.Wand.Wood != "" {
				uniqueSet[char.Wand.Wood] = true
			}
		case "Wand.Core":
			if char.Wand.Core != "" {
				uniqueSet[char.Wand.Core] = true
			}
		case "Actor":
			if char.Actor != "" {
				uniqueSet[char.Actor] = true
			}
		}
	}

	// Convert the set to a slice
	values := []string{}
	for value := range uniqueSet {
		values = append(values, value)
	}
	return values
}

// GenerateOptions ensures unique options by filtering out duplicates
func GenerateOptions(allValues []string, correctAnswer string, rng *rand.Rand) []string {
	wrongOptions := []string{}
	for _, val := range allValues {
		if val != correctAnswer {
			wrongOptions = append(wrongOptions, val)
		}
	}

	// Limit wrong options to 3
	if len(wrongOptions) > 3 {
		wrongOptions = wrongOptions[:3]
	}

	// Add the correct answer
	options := append(wrongOptions, correctAnswer)

	// Shuffle the options
	rng.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	return options
}

// GenerateOptionsFromCharacters now supports AlternateNames
func GenerateOptionsFromCharacters(characters []models.Character, correctValue string, field string, rng *rand.Rand) []string {
	wrongCharacters := FilterCharactersWithDifferentAttribute(characters, field, correctValue)

	// Ensure there are enough wrong options
	if len(wrongCharacters) < 3 {
		return []string{correctValue} // Return only the correct answer
	}

	// Shuffle and select up to 3 incorrect options
	rng.Shuffle(len(wrongCharacters), func(i, j int) {
		wrongCharacters[i], wrongCharacters[j] = wrongCharacters[j], wrongCharacters[i]
	})

	// Extract the field values
	incorrectOptions := []string{}
	for i := 0; i < 3; i++ {
		switch field {
		case "Name":
			incorrectOptions = append(incorrectOptions, wrongCharacters[i].Name)
		case "AlternateNames":
			// Pick a random alternate name from the wrong characters
			if len(wrongCharacters[i].AlternateNames) > 0 {
				incorrectOptions = append(incorrectOptions, wrongCharacters[i].AlternateNames[0])
			}
		}
	}

	// Add the correct answer
	options := append(incorrectOptions, correctValue)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	return options
}
