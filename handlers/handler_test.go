package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"quoter/models"
	"quoter/storage"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

// Setup test router with mock storage
func setupTestRouter() (*mux.Router, *storage.QuoteStorage) {
	storage := storage.NewQuoteStorage()
	handler := NewQuoteHandler(storage)

	router := mux.NewRouter()
	router.HandleFunc("/quotes", handler.GetQuotesHandler).Methods("GET")
	router.HandleFunc("/quotes/random", handler.GetRandomQuoteHandler).Methods("GET")
	router.HandleFunc("/quotes", handler.AddQuoteHandler).Methods("POST")
	router.HandleFunc("/quotes/{id}", handler.DeleteQuoteHandler).Methods("DELETE")

	return router, storage
}

// Test adding a quote
func TestAddQuoteHandler(t *testing.T) {
	router, _ := setupTestRouter()

	body := []byte(`{"author":"Confucius", "quote":"Life is simple."}`)
	req, _ := http.NewRequest("POST", "/quotes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", res.Code)
	}
}

// Test retrieving all quotes
func TestGetQuotesHandler(t *testing.T) {
	router, storage := setupTestRouter()
	storage.AddQuote("Confucius", "Life is simple.")

	req, _ := http.NewRequest("GET", "/quotes", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", res.Code)
	}

	var quotes []models.Quote
	json.Unmarshal(res.Body.Bytes(), &quotes)

	if len(quotes) != 1 {
		t.Errorf("Expected 1 quote, got %d", len(quotes))
	}
}

// Test retrieving a random quote
func TestGetRandomQuoteHandler(t *testing.T) {
	router, storage := setupTestRouter()
	storage.AddQuote("Confucius", "Life is simple.")

	req, _ := http.NewRequest("GET", "/quotes/random", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", res.Code)
	}
}

// Test deleting a quote
func TestDeleteQuoteHandler(t *testing.T) {
	router, storage := setupTestRouter()
	q := storage.AddQuote("Confucius", "Life is simple.")

	req, _ := http.NewRequest("DELETE", "/quotes/"+strconv.Itoa(q.ID), nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		t.Errorf("Expected status NoContent, got %d", res.Code)
	}

	if len(storage.GetQuotes()) != 0 {
		t.Errorf("Expected quotes to be empty after deletion")
	}
}
