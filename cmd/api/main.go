package main

import (
	"flag"
	"log"
	"net/http"

	"go-hexagonal/internal/adapters/handler"
	"go-hexagonal/internal/adapters/repository"
	"go-hexagonal/internal/core/ports"
	"go-hexagonal/internal/core/services"
)

func main() {
	// 0. Define db type to use
	dbType := flag.String("db", "memory", "Database type to use: `memory` or `sqlite`")
	flag.Parse()

	var repo ports.UserRepository
	var err error

	// 1. Driven Adapter
	// Example using an In-Memory Repository
	// repo := repository.NewInMemoryRepo()

	// Example using a SQLite Repository
	switch *dbType {
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