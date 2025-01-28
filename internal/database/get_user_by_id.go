package database

import (
	"Restringing-V2/entity"
	"fmt"

	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func (s *service) GetUserById(id int) (entity.User, error) {
	query := `SELECT * FROM users WHERE id = ?`
	var user entity.User

	// Execute the query
	row := s.db.QueryRow(query, id)

	// Scan the result into the user struct
	err := row.Scan(&user.ID, &user.FirstName, &user.Surname, &user.Username, &user.Email, &user.Password /*, &user.EmailVerification, &user.HasAcceptedTerms*/, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user with id %d not found", id)
		}
		return user, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}
