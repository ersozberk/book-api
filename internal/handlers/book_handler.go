package handlers

import (
	"book-api/internal/models"
	"encoding/json"
	"net/http"
	"sync"
)

var (
	// Veriyi artık models paketinden çekiyoruz
	books = []models.Book{
		{ID: "1", Title: "Go Programming Blueprints", Author: "Mat Ryer", Price: 39.99},
	}
	mu sync.Mutex
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	json.NewEncoder(w).Encode(books)
	mu.Unlock()
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Geçersiz JSON formatı", http.StatusBadRequest)
		return
	}

	mu.Lock()
	books = append(books, newBook)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func BooksRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetBooks(w, r)
	case http.MethodPost:
		AddBook(w, r)
	default:
		http.Error(w, "Bu metod desteklenmiyor", http.StatusMethodNotAllowed)
	}
}
