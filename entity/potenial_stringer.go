package entity

type PotentialStringer struct {
	ID                uint   `json:"id" binding:"required"`
	YearsOfExperience int    `json:"yearsOfExperience" binding:"required"`
	Message           string `json:"message" binding:"required"`
	Status            bool   `json:"status" binding:"required"`
}

var mockPotenialStringer = PotentialStringer{
	ID:                MockUser.ID,
	YearsOfExperience: 3,
	Message:           "Hello my name is adfskl",
	Status:            false,
}
