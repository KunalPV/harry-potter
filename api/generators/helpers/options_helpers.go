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

// GenerateOptionsFromCharacters ensures only one correct answer appears among options
func GenerateOptionsFromCharacters(characters []models.Character, correctName string, field string, rng *rand.Rand) []string {
	// Filter characters with different attribute values
	wrongCharacters := FilterCharactersWithDifferentAttribute(characters, field, correctName)

	// Ensure there are enough wrong options
	if len(wrongCharacters) < 3 {
		// Return a smaller set, but this might require skipping the question
		return []string{correctName} // Only include the correct answer
	}

	// Shuffle and select up to 3 incorrect options
	rng.Shuffle(len(wrongCharacters), func(i, j int) {
		wrongCharacters[i], wrongCharacters[j] = wrongCharacters[j], wrongCharacters[i]
	})
	incorrectNames := []string{wrongCharacters[0].Name, wrongCharacters[1].Name, wrongCharacters[2].Name}

	// Add the correct answer
	options := append(incorrectNames, correctName)

	// Shuffle all options
	rng.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })

	return options
}
