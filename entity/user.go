package entity

import (
	"time"
)

type User struct {
	FirstName string
	Surname   string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
