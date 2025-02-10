package database

import (
	"Restringing-V2/entity"
	"fmt"

	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func (s *service) GetUserByEmail(email string) (entity.User, error) {

	query := `SELECT * FROM users WHERE email = ?`
	var user entity.User

	// Execute the query
	row := s.db.QueryRow(query, email)

	// Scan the result into the user struct
	err := row.Scan(&user.ID, &user.FirstName, &user.Surname, &user.Username, &user.Email, &user.Password /*, &user.EmailVerification, &user.HasAcceptedTerms*/, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with id %d not found", email)
		}
		return user, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil

}
