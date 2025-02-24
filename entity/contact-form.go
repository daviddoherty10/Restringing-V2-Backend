package entity

import "time"

type ContactForm struct {
	Email              string
	Message            string
	HasBeenRespondedTo int
	CreatedAt          time.Time
}
