package helpers

import (
	"fmt"
	"harry-potter-api/models"
)

// FilterCharactersWithAttribute filters characters based on a specific attribute
func FilterCharactersWithAttribute(characters []models.Character, attribute string) []models.Character {
	filtered := []models.Character{}
	for _, char := range characters {
		switch attribute {
		case "Patronus":
			if char.Patronus != "" {
				filtered = append(filtered, char)
			}
		case "House":
			if char.House != "" {
				filtered = append(filtered, char)
			}
		case "Ancestry":
			if char.Ancestry != "" {
				filtered = append(filtered, char)
			}
		case "Wand.Wood":
			if char.Wand.Wood != "" {
				filtered = append(filtered, char)
			}
		case "Wand.Core":
			if char.Wand.Core != "" {
				filtered = append(filtered, char)
			}
		}
	}

	fmt.Printf("Filtered %d characters with attribute %s\n", len(filtered), attribute)
	return filtered
}

// FilterCharactersWithSameAttribute filters characters with the same attribute value
func FilterCharactersWithSameAttribute(characters []models.Character, attribute string, value string) []models.Character {
	filtered := []models.Character{}
	for _, char := range characters {
		switch attribute {
		case "House":
			if char.House == value {
				filtered = append(filtered, char)
			}
		case "Ancestry":
			if char.Ancestry == value {
				filtered = append(filtered, char)
			}
		case "Wand.Wood":
			if char.Wand.Wood == value {
				filtered = append(filtered, char)
			}
			// Add more cases for other attributes as needed
		}
	}
	return filtered
}

// FilterCharactersWithDifferentAttribute filters characters whose attribute value is different
func FilterCharactersWithDifferentAttribute(characters []models.Character, attribute string, value string) []models.Character {
	filtered := []models.Character{}
	for _, char := range characters {
		switch attribute {
		case "Name":
			if char.Name != "" && char.Name != value {
				filtered = append(filtered, char)
			}
		case "House":
			if char.House != "" && char.House != value {
				filtered = append(filtered, char)
			}
		case "Ancestry":
			if char.Ancestry != "" && char.Ancestry != value {
				filtered = append(filtered, char)
			}
		case "Wand.Wood":
			if char.Wand.Wood != "" && char.Wand.Wood != value {
				filtered = append(filtered, char)
			}
		case "Wand.Core":
			if char.Wand.Core != "" && char.Wand.Core != value {
				filtered = append(filtered, char)
			}
		}
	}
	return filtered
}
