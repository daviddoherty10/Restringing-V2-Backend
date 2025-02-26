package entity

import (
	"time"
)

type User struct {
	ID                uint
	FirstName         string
	Surname           string
	Username          string
	Email             string
	EmailVerification bool
	HasAcceptedTerms  bool
	Password          string
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
	HasAcceptedTerms:  true,
	Password:          "password", // NEVER store plain-text passwords in production!
	CreatedAt:         time.Now(),
	UpdatedAt:         time.Now(),
}
