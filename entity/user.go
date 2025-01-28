package entity

import (
	"time"
)

type User struct {
	ID                int
	FirstName         string
	Surname           string
	Username          string
	Email             string
	EmailVerification bool
	Password          string
	HasAcceptedTerms  bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// Mock user for simplicity
var MockUser = User{
	ID:                1,
	FirstName:         "firstname",
	Surname:           "surname",
	Username:          "testuser",
	Email:             "test@testmail.com",
	EmailVerification: true,
	Password:          "password", // NEVER store plain-text passwords in production!
	HasAcceptedTerms:  true,
	CreatedAt:         time.Now(),
	UpdatedAt:         time.Now(),
}
