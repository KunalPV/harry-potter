package utils

import (
	"fmt"
	"math/rand"
	"time"

	"harry-potter-api/models"
)

func GenerateHouseQuestion(characters []models.Character) (map[string]interface{}, error) {
	rand.Seed(time.Now().UnixNano())

	// Select a random character
	character := characters[rand.Intn(len(characters))]

	// Ensure we have valid data
	if character.House == "" {
		return nil, fmt.Errorf("character %s does not have a house assigned", character.Name)
	}

	// Generate options
	houses := extractUniqueHouses(characters)
	options := generateOptions(houses, character.House)

	// Formulate the question
	question := map[string]interface{}{
		"question":   fmt.Sprintf("What is the house of %s?", character.Name),
		"options":    options,
		"answer":     character.House,
		"difficulty": "easy",
	}

	return question, nil
}

// Helper: Extract unique houses
func extractUniqueHouses(characters []models.Character) []string {
	houseSet := make(map[string]bool)
	for _, char := range characters {
		houseSet[char.House] = true
	}

	houses := []string{}
	for house := range houseSet {
		houses = append(houses, house)
	}
	return houses
}

// Helper: Generate options
func generateOptions(allValues []string, correctAnswer string) []string {
	wrongOptions := []string{}
	for _, val := range allValues {
		if val != correctAnswer {
			wrongOptions = append(wrongOptions, val)
		}
	}

	if len(wrongOptions) > 3 {
		wrongOptions = wrongOptions[:3]
	}

	options := append(wrongOptions, correctAnswer)
	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	return options
}
