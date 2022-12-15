package router

import (
	"github.com/briannag31/mv-project3-quiz/middleware"
	"github.com/gorilla/mux"
)

// Future enhancement: delete by collection
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/cards", middleware.GetAllCards).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/decks", middleware.GetAllDecks).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/cards/{id}", middleware.GetOneCard).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/decks/{id}", middleware.GetOneDeck).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cards", middleware.CreateCard).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/decks", middleware.CreateDeck).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/cardKnown/{id}", middleware.CardKnown).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoCardKnown/{id}", middleware.UndoKnown).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteCards/{id}", middleware.DeleteCard).Methods("DELETE", "OPTIONS")
	// router.HandleFunc("/api/deleteCards", middleware.DeleteAllCards).Methods("DELETE", "OPTIONS")

	return router
}
