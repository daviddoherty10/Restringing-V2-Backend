package database

import (
	//"context"
	//"database/sql"
	"fmt"
	//"log"
	//"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func (s *service) getUser1() (User, error) {
	query := `SELECT id, name, email, created_at FROM users`

	var u User
	rows, err := s.db.Query(query)
	if err != nil {
		return u, fmt.Errorf("Unable to get first User", err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return u, fmt.Errorf("Unable to get first User: ", err)
		}
		return u, nil
	}

	return u, fmt.Errorf("Unable to get first User: ", err)

}
