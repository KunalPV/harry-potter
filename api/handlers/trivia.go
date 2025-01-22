package handlers

import (
	"encoding/json"
	"net/http"

	"harry-potter-api/generators"
	"harry-potter-api/models"
)

type TriviaHandler struct {
	Characters []models.Character
	Spells     []models.Spell
}

func NewTriviaHandler(characters []models.Character, spells []models.Spell) *TriviaHandler {
	return &TriviaHandler{
		Characters: characters,
		Spells:     spells,
	}
}

func (h *TriviaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle case where no characters are loaded
	if len(h.Characters) == 0 {
		http.Error(w, "No characters data available", http.StatusInternalServerError)
		return
	}

	// Generate a house-related question
	question, err := generators.GenerateQuestion(h.Characters, h.Spells)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the question
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}
