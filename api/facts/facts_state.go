package facts

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

type Fact struct {
	Fact   string `json:"fact"`
	Source string `json:"source"`
}

type FactState struct {
	CurrentFact    Fact      `json:"current_fact"`
	ShownFacts     []Fact    `json:"shown_facts"`
	LastUpdatedDay time.Time `json:"last_updated_day"`
}

var factsState FactState

// GetCurrentFact returns the current daily fact
func GetCurrentFact() Fact {
	return factsState.CurrentFact
}

// GetLastUpdatedDay returns the last updated day
func GetLastUpdatedDay() time.Time {
	return factsState.LastUpdatedDay
}

// initializeDefaultState sets up the default state
func initializeDefaultState() {
	factsState = FactState{
		ShownFacts:     []Fact{},
		LastUpdatedDay: time.Now().AddDate(0, 0, -1), // Force update on first run
	}
}

// LoadState initializes the state from the state.json file
func LoadState(stateFile string) {
	data, err := os.ReadFile(stateFile)
	if err == nil {
		if len(data) == 0 {
			log.Println("State file is empty. Initializing with default state.")
			initializeDefaultState()
			SaveState(stateFile)
			return
		}
		if err := json.Unmarshal(data, &factsState); err != nil {
			log.Fatalf("Error parsing state file: %v", err)
		}
	} else {
		// Initialize with default values if state file is missing
		log.Println("State file not found. Initializing with default state.")
		initializeDefaultState()
		SaveState(stateFile)
	}
}

// SaveState persists the current state to state.json
func SaveState(stateFile string) {
	data, err := json.MarshalIndent(factsState, "", "  ")
	if err != nil {
		log.Printf("Error saving state: %v", err)
		return
	}
	if err := os.WriteFile(stateFile, data, 0644); err != nil {
		log.Printf("Error writing state file: %v", err)
	}
}

// UpdateDailyFact selects a new fact for the day
func UpdateDailyFact() {
	if len(factsState.ShownFacts) == len(facts) {
		// Reset cycle if all facts have been shown
		factsState.ShownFacts = []Fact{}
	}

	// Select a random fact not already shown
	remainingFacts := []Fact{}
	for _, fact := range facts {
		if !isFactShown(fact) {
			remainingFacts = append(remainingFacts, fact)
		}
	}

	if len(remainingFacts) == 0 {
		return
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	selectedFact := remainingFacts[rand.Intn(len(remainingFacts))]

	// Update the state
	factsState.CurrentFact = selectedFact
	factsState.ShownFacts = append(factsState.ShownFacts, selectedFact)
	factsState.LastUpdatedDay = time.Now().UTC().Truncate(24 * time.Hour)

	// Save the updated state
	SaveState("facts/state/state.json")
}

// isFactShown checks if a fact has already been shown
func isFactShown(fact Fact) bool {
	for _, shownFact := range factsState.ShownFacts {
		if shownFact.Fact == fact.Fact {
			return true
		}
	}
	return false
}
