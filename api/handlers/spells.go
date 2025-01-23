package handlers

import (
	"encoding/json"
	"harry-potter-api/models"
	"net/http"
	"strconv"
)

// GetSpellsHandler serves paginated spell data
func GetSpellsHandler(w http.ResponseWriter, r *http.Request, spells []models.Spell) {
	// Parse pagination query parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 30
	}

	// Calculate pagination indices
	start := (page - 1) * limit
	end := start + limit

	// Handle out-of-bounds indices
	if start >= len(spells) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]models.Spell{}) // Return empty list for out-of-bounds page
		return
	}
	if end > len(spells) {
		end = len(spells)
	}

	// Fetch paginated spells
	paginatedSpells := spells[start:end]

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":       paginatedSpells,
		"page":       page,
		"limit":      limit,
		"total":      len(spells),
		"totalPages": (len(spells) + limit - 1) / limit, // Ceiling division
	})
}
