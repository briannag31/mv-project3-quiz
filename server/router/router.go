package router

import (
	"github.com/briannag31/mv-project3-quiz/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/cards", middleware.GetAllCards).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cards/{id}", middleware.GetOneCard).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/cards", middleware.CreateCard).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/cards/{id}", middleware.CardKnown).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/cards/{id}", middleware.UndoKnown).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/cards/{id}", middleware.DeleteCard).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/cards", middleware.DeleteAllCards).Methods("DELETE", "OPTIONS")

	return router
}