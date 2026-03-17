package main

import (
	"book-api/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	// Rotaları (Routes) handlers paketinden alarak bağlıyoruz
	http.HandleFunc("/api/books", handlers.BooksRouter)

	port := ":8080"
	fmt.Printf("🚀 Modüler API Sunucusu %s portunda dinleniyor...\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Sunucu başlatılırken hata:", err)
	}
}
