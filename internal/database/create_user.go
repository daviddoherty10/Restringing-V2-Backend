package database

import (
	"Restringing-V2/entity"
	"fmt"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func (s *service) CreateUser(u entity.User) error {
	// Define the SQL INSERT query
	query := `INSERT INTO users (firstname, surname, username, email, email_confirmed, has_accepted_terms, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the query with the provided parameters
	_, err := s.db.Exec(query, u.FirstName, u.Surname, u.Username, u.Email, u.EmailVerification, u.HasAcceptedTerms, u.Password, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	log.Printf("User created successfully: Name=%s, Email=%s", u.FirstName, u.Email)
	return nil
}
