package storage

import (
	"testing"
)

// Test adding a quote
func TestAddQuote(t *testing.T) {
	s := NewQuoteStorage()
	q := s.AddQuote("Confucius", "Life is simple.")

	if q.ID != 1 {
		t.Errorf("Expected ID 1, got %d", q.ID)
	}

	if q.Author != "Confucius" {
		t.Errorf("Expected author Confucius, got %s", q.Author)
	}

	if q.Text != "Life is simple." {
		t.Errorf("Expected quote text, got %s", q.Text)
	}
}

// Test retrieving all quotes
func TestGetQuotes(t *testing.T) {
	s := NewQuoteStorage()
	s.AddQuote("Confucius", "Life is simple.")
	s.AddQuote("Aristotle", "Knowing yourself is the beginning of all wisdom.")

	quotes := s.GetQuotes()
	if len(quotes) != 2 {
		t.Errorf("Expected 2 quotes, got %d", len(quotes))
	}
}

// Test retrieving a random quote
func TestGetRandomQuote(t *testing.T) {
	s := NewQuoteStorage()
	s.AddQuote("Confucius", "Life is simple.")

	q, found := s.GetRandomQuote()
	if !found {
		t.Errorf("Expected a quote, got nothing")
	}

	if q.Author != "Confucius" {
		t.Errorf("Expected Confucius, got %s", q.Author)
	}
}

// Test deleting a quote
func TestDeleteQuote(t *testing.T) {
	s := NewQuoteStorage()
	q := s.AddQuote("Confucius", "Life is simple.")

	if !s.DeleteQuote(q.ID) {
		t.Errorf("Expected deletion to succeed")
	}

	if len(s.GetQuotes()) != 0 {
		t.Errorf("Expected empty storage after deletion")
	}
}
