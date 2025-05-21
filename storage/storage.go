package storage

import (
	"math/rand"
	"quoter/models"
	"sync"
)

// QuoteStorage manages in-memory storage
type QuoteStorage struct {
	quotes []models.Quote
	nextID int
	mu     sync.Mutex
}

// NewQuoteStorage initializes storage
func NewQuoteStorage() *QuoteStorage {
	return &QuoteStorage{quotes: []models.Quote{}, nextID: 1}
}

// AddQuote stores a new quote
func (s *QuoteStorage) AddQuote(author, text string) models.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	newQuote := models.Quote{ID: s.nextID, Author: author, Text: text}
	s.quotes = append(s.quotes, newQuote)
	s.nextID++
	return newQuote
}

// GetQuotes returns all quotes
func (s *QuoteStorage) GetQuotes() []models.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.quotes
}

// GetRandomQuote returns a random quote
func (s *QuoteStorage) GetRandomQuote() (models.Quote, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.quotes) == 0 {
		return models.Quote{}, false
	}
	return s.quotes[rand.Intn(len(s.quotes))], true
}

// DeleteQuote removes a quote by ID
func (s *QuoteStorage) DeleteQuote(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, q := range s.quotes {
		if q.ID == id {
			s.quotes = append(s.quotes[:i], s.quotes[i+1:]...)
			return true
		}
	}
	return false
}
