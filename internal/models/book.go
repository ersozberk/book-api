package models

// Book, sistemimizdeki kitap yapısını temsil eder
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
