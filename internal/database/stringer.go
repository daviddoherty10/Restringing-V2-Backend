package database

import (
	"Restringing-V2/entity"
	"errors"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func (s *service) CreatePotenialStringer(p entity.PotentialStringer) error {
	// Define the SQL INSERT query
	query := `INSERT INTO potenialStringer (id, yearsOfExperience,message,status) VALUES (?, ?, ?, ?, ?, ?, ?)`

	// Execute the query with the provided parameters
	_, err := s.db.Exec(query, p.ID, p.YearsOfExperience, p.Message, p.Status)
	if err != nil {
		return errors.New("Failed to create Potenial Stringer: " + err.Error())
	}

	return nil
}
