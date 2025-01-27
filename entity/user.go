package entity

import (
	"time"
)

type User struct {
	ID               int
	FirstName        string
	Surname          string
	Email            string
	Password         string
	HasAcceptedTerms bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
