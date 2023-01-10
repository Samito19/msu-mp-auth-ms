package models

import "github.com/google/uuid"

type User struct {
	Uuid         uuid.UUID
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}
