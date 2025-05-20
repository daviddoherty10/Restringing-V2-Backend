package database

import (
	"Restringing-V2/entity"
	"fmt"

	"database/sql"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func (s *service) GetUserByUsername(username string) (entity.User, error) {

	query := `SELECT * FROM users WHERE username = ?`
	var user entity.User

	// Execute the querymulti
	row := s.db.QueryRow(query, username)

	// Scan the result into the user struct
	err := row.Scan(&user.ID, &user.FirstName, &user.Surname, &user.Username, &user.Email, &user.EmailVerification, &user.HasAcceptedTerms, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with id %d not found", username)
		}
		return user, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil

}

func (s *service) GetUserById(id uint) (entity.User, error) {
	query := `SELECT * FROM users WHERE id = ?`
	var user entity.User

	// Execute the query
	row := s.db.QueryRow(query, id)

	// Scan the result into the user struct
	err := row.Scan(&user.ID, &user.FirstName, &user.Surname, &user.Username, &user.Email, &user.EmailVerification, &user.HasAcceptedTerms, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with id %d not found", id)
		}
		return user, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}

func (s *service) GetUserByEmail(email string) (entity.User, error) {

	query := `SELECT * FROM users WHERE email = ?`
	var user entity.User

	// Execute the query
	row := s.db.QueryRow(query, email)

	// Scan the result into the user struct
	err := row.Scan(&user.ID, &user.FirstName, &user.Surname, &user.Username, &user.Email, &user.EmailVerification, &user.HasAcceptedTerms, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("user with id %d not found", email)
			return user, fmt.Errorf("user with id %d not found", email)
		}
		log.Println("failed to get user: %v", err)
		return user, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil

}

func (s *service) DeleteUser(id uint) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateUser(u entity.User) error {
	// Define the SQL INSERT query
	query := `INSERT INTO users (firstname, surname, username, email, email_confirmed, has_accepted_terms, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the query with the provided parameters
	_, err := s.db.Exec(query, u.FirstName, u.Surname, u.Username, u.Email, u.EmailVerification, u.HasAcceptedTerms, u.Password, time.Now(), time.Now())
	if err != nil {
		log.Println("failed to insert user: %w", err)
		return fmt.Errorf("failed to insert user: %w", err)
	}

	log.Printf("User created successfully: Name=%s, Email=%s", u.FirstName, u.Email)
	return nil
}

func (s *service) UpdateUser(u entity.User) error {
	query := `UPDATE users (firstname, surname, username, email, email_confirmed, has_accepted_terms, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) WHRE id = ? VALUES (?)`

	_, err := s.db.Exec(query, u.FirstName, u.Surname, u.Username, u.Email, u.EmailVerification, u.HasAcceptedTerms, u.Password, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
