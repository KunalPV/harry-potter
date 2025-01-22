package main

import (
	"fmt"
	"harry-potter-api/handlers"
	"harry-potter-api/models"
	"harry-potter-api/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var characters []models.Character
var spells []models.Spell

func main() {
	// Load JSON data
	utils.LoadJSON("data/characters.json", &characters)
	utils.LoadJSON("data/spells.json", &spells)

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

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
