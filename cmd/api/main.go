package main

import (
	"log"
	"net/http"

	"go-hexagonal/internal/adapters/handler"
	"go-hexagonal/internal/adapters/repository"
	"go-hexagonal/internal/core/services"
)

func main() {
	// 1. Driven Adapter
	// Example using an In-Memory Repository
	// repo := repository.NewInMemoryRepo()

	// Example using a SQLite Repository
	repo, err := repository.NewSQliteRepository("./users.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 2. Core Service
	service := services.NewUserService(repo)

	// Driver Adapter HTTP Handler
	handler := handler.NewUserHandler(service)

	// 4. Router (Go standard lib)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", handler.SaveUser)
	mux.HandleFunc("GET /users/{nickname}", handler.GetUser)

	// 4. Logs
	log.Println("Starting server at port 8080")

	// 5. Server
	log.Fatal(http.ListenAndServe(":8080", mux))
}