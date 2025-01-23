package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"harry-potter-api/facts"
)

var (
	stateMutex sync.Mutex
)

// GetFactHandler serves the current fact
func GetFactHandler(w http.ResponseWriter, r *http.Request) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	today := time.Now().UTC().Truncate(24 * time.Hour)
	if facts.GetLastUpdatedDay().Before(today) {
		// Update fact if a new day has started
		facts.UpdateDailyFact()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(facts.GetCurrentFact())
}
