package handlers

import (
	"encoding/json"
	"net/http"
)

func TriviaHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Trivia question endpoint is under construction.",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
