package utils

import (
	"fmt"
	"math/rand"
	"time"

	"harry-potter-api/models"
)

// GenerateQuestion generates a random trivia question
func GenerateQuestion(characters []models.Character, spells []models.Spell) (map[string]interface{}, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Question types
	questionTypes := []string{"Name", "House", "Ancestry", "Wand Wood", "Wand Core", "Patronus", "Hair Color", "Student/Staff", "Actor", "Spell Name", "Spell Description"}
	selectedType := questionTypes[rng.Intn(len(questionTypes))]

	switch selectedType {
	// case "Name":
	// 	return generateNameQuestion(characters, rng)
	case "House":
		return GenerateHouseQuestion(characters)
	// case "Ancestry":
	// 	return generateAncestryQuestion(characters, rng)
	// case "Wand Wood":
	// 	return generateWandWoodQuestion(characters, rng)
	// case "Wand Core":
	// 	return generateWandCoreQuestion(characters, rng)
	// case "Patronus":
	// 	return generatePatronusQuestion(characters, rng)
	// case "Hair Color":
	// 	return generateHairColorQuestion(characters, rng)
	// case "Student/Staff":
	// 	return generateStudentStaffQuestion(characters, rng)
	// case "Actor":
	// 	return generateActorQuestion(characters, rng)
	// case "Spell Name":
	// 	return generateSpellNameQuestion(spells, rng)
	// case "Spell Description":
	// 	return generateSpellDescriptionQuestion(spells, rng)
	default:
		return nil, fmt.Errorf("unknown question type: %s", selectedType)
	}
}

func GenerateHouseQuestion(characters []models.Character) (map[string]interface{}, error) {
	// Filter out characters without a house
	validCharacters := []models.Character{}
	for _, char := range characters {
		if char.House != "" {
			validCharacters = append(validCharacters, char)
		}
	}

	// Check if there are valid characters left
	if len(validCharacters) == 0 {
		return nil, fmt.Errorf("no characters with a house available to generate questions")
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	character := validCharacters[rng.Intn(len(validCharacters))]

	// Generate options
	houses := extractUniqueHouses(validCharacters)
	if len(houses) < 4 {
		return nil, fmt.Errorf("not enough unique houses to generate options")
	}

	options := generateOptions(houses, character.House)

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
		if char.House != "" {
			houseSet[char.House] = true
		}
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
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	return options
}
