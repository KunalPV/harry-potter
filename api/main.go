package main

import (
	"fmt"
	"harry-potter-api/facts"
	"harry-potter-api/handlers"
	"harry-potter-api/models"
	"harry-potter-api/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "harry-potter-api/generators/actor"
	_ "harry-potter-api/generators/alternate_names"
	_ "harry-potter-api/generators/ancestry"
	_ "harry-potter-api/generators/house"
	_ "harry-potter-api/generators/name"
	_ "harry-potter-api/generators/spell_description"
	_ "harry-potter-api/generators/spell_name"
	_ "harry-potter-api/generators/wand_core"
	_ "harry-potter-api/generators/wand_wood"
)

var characters []models.Character
var spells []models.Spell

func main() {
	// Load JSON data
	utils.LoadJSON("data/characters.json", &characters)
	utils.LoadJSON("data/spells.json", &spells)

	// Initialize facts and state
	facts.LoadFacts("data/facts.json")
	facts.LoadState("facts/state/state.json")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Initialize TriviaHandler
	triviaHandler := handlers.NewTriviaHandler(characters, spells)

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Harry Potter Trivia API!"))
	})

	r.Get("/api/game/question/", triviaHandler.ServeHTTP)
	r.Get("/api/facts", handlers.GetFactHandler)
	r.Get("/api/characters", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCharactersHandler(w, r, characters)
	})
	r.Get("/api/spells", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSpellsHandler(w, r, spells)
	})

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
