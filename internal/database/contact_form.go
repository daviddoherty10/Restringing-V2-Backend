package database

import (
	"Restringing-V2/entity"
	"errors"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func (s *service) CreateContactFormResponse(res entity.ContactForm) error {
	query := `INSERT INTO contact_form (email,message,has_been_responded_to,created_at) VALUES (?, ?, ?,?)`
	_, err := s.db.Exec(query, res.Email, res.Message, 0, time.Now())
	if err != nil {
		return errors.New("Unable to Create Contact Form response" + err.Error())
	}

	return nil
}
