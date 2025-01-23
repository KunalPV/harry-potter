package handlers

import (
	"encoding/json"
	"harry-potter-api/models"
	"net/http"
	"strconv"
)

// GetCharactersHandler serves paginated character data
func GetCharactersHandler(w http.ResponseWriter, r *http.Request, characters []models.Character) {
	// Parse pagination query parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 20
	}

	// Calculate pagination indices
	start := (page - 1) * limit
	end := start + limit

	// Handle out-of-bounds indices
	if start >= len(characters) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]models.Character{}) // Return empty list for out-of-bounds page
		return
	}
	if end > len(characters) {
		end = len(characters)
	}

	// Fetch paginated characters
	paginatedCharacters := characters[start:end]

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":       paginatedCharacters,
		"page":       page,
		"limit":      limit,
		"total":      len(characters),
		"totalPages": (len(characters) + limit - 1) / limit, // Ceiling division
	})
}
