package handlers

import (
	"encoding/json"
	"net/http"
	"quoter/models"
	"quoter/storage"
	"strconv"

	"github.com/gorilla/mux"
)

// QuoteHandler wraps storage for request handling
type QuoteHandler struct {
	Storage *storage.QuoteStorage
}

// NewQuoteHandler initializes handlers with storage dependency
func NewQuoteHandler(s *storage.QuoteStorage) *QuoteHandler {
	return &QuoteHandler{Storage: s}
}

// GetQuotesHandler fetches all quotes (supports author filtering)
func (h *QuoteHandler) GetQuotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authorFilter := r.URL.Query().Get("author")
	quotes := h.Storage.GetQuotes()

	if authorFilter != "" {
		filteredQuotes := []models.Quote{}
		for _, q := range quotes {
			if q.Author == authorFilter {
				filteredQuotes = append(filteredQuotes, q)
			}
		}
		json.NewEncoder(w).Encode(filteredQuotes)
		return
	}

	json.NewEncoder(w).Encode(quotes)
}

// GetRandomQuoteHandler fetches a random quote
func (h *QuoteHandler) GetRandomQuoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	quote, found := h.Storage.GetRandomQuote()
	if !found {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(quote)
}

// AddQuoteHandler adds a new quote
func (h *QuoteHandler) AddQuoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newQuote models.Quote
	if err := json.NewDecoder(r.Body).Decode(&newQuote); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newQuote = h.Storage.AddQuote(newQuote.Author, newQuote.Text)
	json.NewEncoder(w).Encode(newQuote)
}

// DeleteQuoteHandler removes a quote by ID
func (h *QuoteHandler) DeleteQuoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if h.Storage.DeleteQuote(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Quote not found", http.StatusNotFound)
	}
}
