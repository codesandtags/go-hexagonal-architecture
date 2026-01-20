package repository

import (
	"database/sql"
	"errors"
	"go-hexagonal/internal/core/domain"

	_ "github.com/mattn/go-sqlite3"
)

type SQliteRepository struct {
	db *sql.DB
}

/**
 * Connection to the dabase and create the table if it doesn't exist
 **/
func NewSQliteRepository(dbPath string) (*SQliteRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	// Check the connection with a Ping
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Simple migration On-the-fly
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		nickname TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &SQliteRepository{db: db}, nil
}

func (r *SQliteRepository) Save(user domain.User) error {
	query := "INSERT INTO users (id, name, email, nickname) VALUES (?, ?, ?, ?)"

	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Nickname)
	if err != nil {
		// Aquí podrías chequear si es error de constraint (UNIQUE) y retornar un error de dominio específico
		return err
	}

	return nil
}

func (r *SQliteRepository) GetByNickname(nickname string) (domain.User, error) {
	query := "SELECT id, name, email, nickname FROM users WHERE nickname = ?"

	row := r.db.QueryRow(query, nickname)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}