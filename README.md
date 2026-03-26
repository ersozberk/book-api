# 📚 Go REST API: Clean Architecture

[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/ersozberk/book-api/edit/main/README.md)
[![pt-br](https://img.shields.io/badge/lang-tr-green.svg)](https://github.com/ersozberk/book-api/edit/main/README-tr.md)

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Architecture](https://img.shields.io/badge/Architecture-Modular%20%2F%20Internal-orange)
![Dependencies](https://img.shields.io/badge/Dependencies-Zero-brightgreen)

A modular RESTful API built entirely from scratch using Go's standard library (`net/http`). This project demonstrates how to structure a scalable Go application by separating concerns into distinct layers without relying on external web frameworks.

## ✨ Features

* **Zero Dependencies:** Pure Go standard library implementation for routing and HTTP handling.
* **Clean Architecture:** Organized using the industry-standard `cmd/` and `internal/` directory layout.
* **Concurrency Safe:** Utilizes `sync.Mutex` to prevent race conditions during concurrent read/write operations on the in-memory data store.
* **RESTful Principles:** Strict adherence to HTTP methods (GET, POST) and proper status codes (200 OK, 201 Created, 400 Bad Request).

## 📂 Project Structure

```text
.
├── cmd/
│   └── server/
│       └── main.go       # Application entry point and route registration
├── internal/
│   ├── handlers/
│   │   └── book.go       # HTTP handlers and business logic
│   └── models/
│       └── book.go       # Data structures and domain models
├── go.mod                # Go module definitions
├── .gitignore            # Git ignore rules
└── README.md             # Project documentation
```

## 🛠 Installation and Usage
Clone the repository:

```bash
git clone [https://github.com/ersozberk/book-api.git](https://github.com/ersozberk/book-api.git)
cd book-api
```

Run the API server:

```bash
go run cmd/server/main.go

```
Test the endpoints (in a separate terminal):

```bash
# Fetch all books (GET)
curl -s http://localhost:8080/api/books

# Add a new book (POST)
curl -s -X POST http://localhost:8080/api/books \
-H "Content-Type: application/json" \
-d '{"id": "2", "title": "Clean Architecture", "author": "Robert C. Martin", "price": 45.00}'
```

Build for production:

```bash
go build -o book-server cmd/server/main.go
./book-server
```
