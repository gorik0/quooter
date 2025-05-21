package router

import (
	"quoter/handlers"
	"quoter/storage"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router with handlers
func SetupRouter(storage *storage.QuoteStorage) *mux.Router {
	handler := handlers.NewQuoteHandler(storage)

	router := mux.NewRouter()
	router.HandleFunc("/quotes", handler.GetQuotesHandler).Methods("GET")
	router.HandleFunc("/quotes/random", handler.GetRandomQuoteHandler).Methods("GET")
	router.HandleFunc("/quotes", handler.AddQuoteHandler).Methods("POST")
	router.HandleFunc("/quotes/{id}", handler.DeleteQuoteHandler).Methods("DELETE")

	return router
}
