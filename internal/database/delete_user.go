package database

import (
	//"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func (s *service) DeleteUser(id uint) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
