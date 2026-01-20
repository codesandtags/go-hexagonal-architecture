package main

import (
	"log"
	"net/http"
	"os"

	"go-hexagonal/internal/adapters/handler"
	"go-hexagonal/internal/adapters/repository"
	"go-hexagonal/internal/core/ports"
	"go-hexagonal/internal/core/services"

	"github.com/joho/godotenv"
)

func main() {

	// Try to Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Using system environment variables")
	}

	// Define db type to use
	dbType := os.Getenv("DB_TYPE")

	// Defult value
	if dbType == "" {
		dbType = "memory"
	}

	var repo ports.UserRepository

	// 1. Selection of Strategy (Strategy Pattern)
	switch dbType {
	case "memory":
		repo = repository.NewInMemoryRepo()
		log.Println("Strategy: In-Memory Storage")
	case "sqlite":
		repo, err = repository.NewSQliteRepository("./users.db")
		if err != nil {
			log.Fatal("Failed to initialize SQLite:", err)
		}
		log.Println("Strategy: SQLite Storage")
	default:
		log.Println("Opción de base de datos no válida. Usando 'memory' por defecto.")
		repo = repository.NewInMemoryRepo()
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