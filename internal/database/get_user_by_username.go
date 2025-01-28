package database

import (
	"Restringing-V2/entity"
	"fmt"

	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func (s *service) GetUserByUsername(username string) (entity.User, error) {

	query := `SELECT * FROM users WHERE username = ?`
	var user entity.User

	// Execute the query
	row := s.db.QueryRow(query, username)

	// Scan the result into the user struct
	err := row.Scan(&user.ID, &user.FirstName, &user.Surname, &user.Username, &user.Email, &user.Password /*, &user.EmailVerification, &user.HasAcceptedTerms*/, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with id %d not found", username)
		}
		return user, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil

}
